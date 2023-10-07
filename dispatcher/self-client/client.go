package main

import (
	"github.com/deyvidm/sms/dispatcher/log"
	"github.com/deyvidm/sms/dispatcher/tasks"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func main() {
	logger = log.GetLogger()
	logger.Info("Test Client initiated")
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	var t *asynq.Task

	t, _ = tasks.NewInviteTask("inviteID123", "+11234567890", "tiger butts")
	logAndEnqueue(client, t)

	// t, _ = tasks.NewReponseTask("+11234567890", "YOOO")
	// logAndEnqueue(client, t)

}

func logAndEnqueue(c *asynq.Client, t *asynq.Task) {
	info, err := c.Enqueue(t)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("Successfully enqueued task %s with ID %v ", info.Type, info.ID)
}
