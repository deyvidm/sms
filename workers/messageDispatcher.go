package workers

import (
	"context"
	"encoding/json"

	"github.com/deyvidm/sms-asynq/client"
	"github.com/deyvidm/sms-asynq/log"
	"github.com/deyvidm/sms-asynq/tasks"
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
	logger.Infof("Received task %s with ID %s", t.Type(), t.ResultWriter().TaskID())
	var p tasks.NewInvitePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	if err := md.irs.SaveNewInviteEntry(p.ToPhoneNumber, p.InviteID); err != nil {
		return err
	}

	// invites, err := md.fetchAllInvites(p.ToPhoneNumber)
	// if err != nil {
	// 	return err
	// }

	// if len(invites) > 1

	// md.fetchContactInvites(p.ToPhoneNumber)
	// return nil
	logger.Info("Sending invite '%s'", p.InviteID)
	// return md.wbc.UpdateInvite(&client.UpdateInvite{
	// 	ID:     p.InviteID,
	// 	Status: utils.Ptr(types.InviteStatus_Invited),
	// })

	return nil
}
