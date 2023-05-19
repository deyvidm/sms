package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const (
	TypeNewResponse = "response:new"
)

type NewResponsePayload struct {
	From    string
	Content string
}

func NewReponseTask(FromPhoneNumber, content string) (*asynq.Task, error) {
	payload, err := json.Marshal(NewResponsePayload{From: FromPhoneNumber, Content: content})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeNewResponse, payload), nil
}
