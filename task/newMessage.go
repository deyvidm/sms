package task

import (
	"context"
	"encoding/json"

	"github.com/deyvidm/sms-asynq/client"
	"github.com/deyvidm/sms-asynq/log"
	"github.com/hibiken/asynq"
)

const (
	TypeNewMessage = "message:new"
)

var logger = log.GetLogger()

type NewMessagePayload struct {
	InviteID      string
	ToPhoneNumber string
	Content       string
}

func NewNewMessageTask(inviteID, toPhoneNumber, content string) (*asynq.Task, error) {
	payload, err := json.Marshal(NewMessagePayload{InviteID: inviteID, ToPhoneNumber: toPhoneNumber, Content: content})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeNewMessage, payload), nil
}

type MessageDispatcher struct {
	wbc client.WebBackendClient
}

func NewMessageDispatcher(wbc client.WebBackendClient) *MessageDispatcher {
	return &MessageDispatcher{
		wbc: wbc,
	}
}

func Ptr[T any](v T) *T {
	return &v
}

func (md *MessageDispatcher) HandleNewMessageTask(ctx context.Context, t *asynq.Task) error {
	logger.Infof("Received task %s with ID %s", t.Type(), t.ResultWriter().TaskID())
	var p NewMessagePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	logger.Printf("Sending message '%s' to '%s'", p.Content, p.ToPhoneNumber)
	md.wbc.UpdateInvite(&client.UpdateInvite{
		ID:     p.InviteID,
		Status: Ptr("invited"),
	})
	return nil
}
