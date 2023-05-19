package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypeNewInvite = "invite:new"
)

type NewInvitePayload struct {
	InviteID      string
	ToPhoneNumber string
	Content       string
}

func NewInviteTask(inviteID string, toPhoneNumber, content string) (*asynq.Task, error) {
	payload, err := json.Marshal(NewInvitePayload{InviteID: inviteID, ToPhoneNumber: toPhoneNumber, Content: content})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeNewInvite, payload), nil
}
