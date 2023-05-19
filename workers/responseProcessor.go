package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/deyvidm/sms-asynq/client"
	"github.com/deyvidm/sms-asynq/tasks"
	"github.com/deyvidm/sms-asynq/types"
	"github.com/deyvidm/sms-asynq/utils"
	"github.com/hibiken/asynq"
)

type ResponseProcessor struct {
	wbc *client.WebBackendClient
	irs *client.InviteResponseStore
}

func NewResponseProcessor(wbc *client.WebBackendClient, irs *client.InviteResponseStore) *ResponseProcessor {
	return &ResponseProcessor{
		wbc: wbc,
		irs: irs,
	}
}

var AcceptedResponseStrings = []string{"yes", "y"}
var DeclinedResponseStrings = []string{"no", "n"}

func (rp *ResponseProcessor) getInviteStatus(content string) string {
	content = strings.ToLower(content)
	if utils.Contains(AcceptedResponseStrings, content) {
		return types.InviteStatus_Accepted
	}
	if utils.Contains(DeclinedResponseStrings, content) {
		return types.InviteStatus_Declined
	}
	return ""
}

func (rp *ResponseProcessor) isValidResponseString(content string) bool {
	content = strings.ToLower(content)
	return utils.Contains(AcceptedResponseStrings, content) ||
		utils.Contains(DeclinedResponseStrings, content)
}

func (rp *ResponseProcessor) HandleResponse(ctx context.Context, t *asynq.Task) error {
	var p tasks.NewResponsePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	logger.Infof("Response from %s : '%s'", p.From, p.Content)

	invites, err := rp.irs.FetchAllInvites(p.From)
	if err != nil {
		return err
	}

	if !rp.isValidResponseString(p.Content) {
		return fmt.Errorf("invalid response '%s'", p.Content)
	}
	if len(invites) == 1 {
		for _, inv := range invites {
			return rp.wbc.UpdateInvite(client.UpdateInvite{
				ID:     inv,
				Status: utils.Ptr(rp.getInviteStatus(p.Content)),
			})
		}
	} else {
		//TODO handle multiple invites
		return nil
	}

	return nil
}
