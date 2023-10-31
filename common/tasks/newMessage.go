package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypeNewMessage = "message:new"
)

type NewMessagePayload struct {
	ToPhoneNumber string
	Content       string
}

func NewMesssageTask(toPhoneNumber, content string) (*asynq.Task, error) {
	payload, err := json.Marshal(NewMessagePayload{ToPhoneNumber: toPhoneNumber, Content: content})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeNewMessage, payload), nil
}
