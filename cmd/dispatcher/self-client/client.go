package main

import (
	"github.com/deyvidm/sms/cmd/dispatcher/log"
	"github.com/deyvidm/sms/pkg/tasks"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func main() {
	logger = log.GetLogger()
	logger.Info("Test Client initiated")
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	var t *asynq.Task

	// t, _ = tasks.NewInviteTask("inviteID123", "+11234567890", "tiger butts")
	t, _ = tasks.NewMesssageTask("+11234567890", "message body goes here")
	logAndEnqueue(client, t)

}

func logAndEnqueue(c *asynq.Client, t *asynq.Task) {
	info, err := c.Enqueue(t)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("Successfully enqueued task %s with ID %v ", info.Type, info.ID)
}
