package workers

import (
	"context"
	"encoding/json"

	"github.com/deyvidm/sms-asynq/client"
	"github.com/deyvidm/sms-asynq/log"
	"github.com/deyvidm/sms-asynq/tasks"
	"github.com/deyvidm/sms-asynq/types"
	"github.com/deyvidm/sms-asynq/utils"
	"github.com/hibiken/asynq"
)

var logger = log.GetLogger()

type MessageDispatcher struct {
	wbc client.WebBackendClient
}

func NewMessageDispatcher(wbc client.WebBackendClient) *MessageDispatcher {
	return &MessageDispatcher{
		wbc: wbc,
	}
}

func (md *MessageDispatcher) HandleNewMessageTask(ctx context.Context, t *asynq.Task) error {
	logger.Infof("Received task %s with ID %s", t.Type(), t.ResultWriter().TaskID())
	var p tasks.NewMessagePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	logger.Info("Sending invite '%s'", p.InviteID)
	return md.wbc.UpdateInvite(&client.UpdateInvite{
		ID:     p.InviteID,
		Status: utils.Ptr(types.InviteStatus_Invited),
	})
}
