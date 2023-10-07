package workers

import (
	"context"
	"encoding/json"

	"github.com/deyvidm/sms/dispatcher/client"
	"github.com/deyvidm/sms/dispatcher/log"
	"github.com/deyvidm/sms/dispatcher/tasks"
	"github.com/deyvidm/sms/dispatcher/types"
	"github.com/deyvidm/sms/dispatcher/utils"
	"github.com/hibiken/asynq"
)

var logger = log.GetLogger()

type MessageDispatcher struct {
	wbc *client.WebBackendClient
	irs *client.InviteResponseStore
}

func NewMessageDispatcher(wbc *client.WebBackendClient, irs *client.InviteResponseStore) *MessageDispatcher {
	return &MessageDispatcher{
		wbc: wbc,
		irs: irs,
	}
}

func (md *MessageDispatcher) HandleSendInviteTask(ctx context.Context, t *asynq.Task) error {
	var p tasks.NewInvitePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	logger.Infof("|%s|\tinviting %s to %s : '%s'", t.Type(), p.ToPhoneNumber, p.InviteID, p.Content)

	if err := md.irs.SaveNewInviteEntry(p.ToPhoneNumber, p.InviteID); err != nil {
		return err
	}

	return md.wbc.UpdateInvite(client.UpdateInvite{
		ID:     p.InviteID,
		Status: utils.Ptr(types.InviteStatus_Invited.String()),
	})
}
