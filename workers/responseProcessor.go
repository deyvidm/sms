package workers

import (
	"context"
	"encoding/json"

	"github.com/deyvidm/sms-asynq/client"
	"github.com/deyvidm/sms-asynq/tasks"
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

func (rp *ResponseProcessor) HandleResponse(ctx context.Context, t *asynq.Task) error {
	logger.Infof("Received task %s with ID %s", t.Type(), t.ResultWriter().TaskID())
	var p tasks.NewResponsePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	invites, err := rp.irs.FetchAllInvites(p.From)
	if err != nil {
		return err
	}

	if len(invites) == 1 {
		// ...
	}

	logger.Infof("Received Response '%s' from '%s'", p.Content, p.From)
	return nil
}
