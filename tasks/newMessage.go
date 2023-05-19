package tasks

import (
	"encoding/json"

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

func NewNewMessageTask(inviteID string, toPhoneNumber, content string) (*asynq.Task, error) {
	payload, err := json.Marshal(NewMessagePayload{InviteID: inviteID, ToPhoneNumber: toPhoneNumber, Content: content})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeNewMessage, payload), nil
}
