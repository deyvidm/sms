package main

import (
	"github.com/deyvidm/sms-asynq/log"
	"github.com/deyvidm/sms-asynq/task"
	"github.com/hibiken/asynq"
)

func main() {
	logger := log.GetLogger()
	logger.Info("Test Client initiated")
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	t1, err := task.NewNewMessageTask("+11234567890", "tiger butts")
	if err != nil {
		logger.Fatal(err)
	}

	logger.Infof("Enquing task %s", t1.Type())
	info, err := client.Enqueue(t1)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("Successfully enqueued task %s with ID %v ", info.Type, info.ID)
}
