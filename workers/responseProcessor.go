package workers

import (
	"context"
	"encoding/json"

	"github.com/deyvidm/sms-asynq/client"
	"github.com/deyvidm/sms-asynq/tasks"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

type ResponseProcessor struct {
	wbc *client.WebBackendClient
	rdb *redis.Client
}

func NewResponseProcessor(wbc *client.WebBackendClient, rdb *redis.Client) *ResponseProcessor {
	return &ResponseProcessor{
		wbc: wbc,
		rdb: rdb,
	}
}

func (rp *ResponseProcessor) HandleResponse(ctx context.Context, t *asynq.Task) error {
	logger.Infof("Received task %s with ID %s", t.Type(), t.ResultWriter().TaskID())
	var p tasks.NewResponsePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	logger.Infof("Received Response '%s' from '%s'", p.Content, p.From)
	return nil
}
