package main

import (
	"github.com/deyvidm/sms-asynq/log"
	"github.com/deyvidm/sms-asynq/tasks"
	"github.com/hibiken/asynq"
)

func main() {
	logger := log.GetLogger()
	logger.Info("Test Client initiated")
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	var err error
	var info *asynq.TaskInfo
	t1, err := tasks.NewInviteTask("123", "+11234567890", "tiger butts")
	if err != nil {
		logger.Fatal(err)
	}

	// logger.Infof("Enquing Invite %s", t1.Type())
	// info, err = client.Enqueue(t1)
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// logger.Infof("Successfully enqueued task %s with ID %v ", info.Type, info.ID)

	logger.Infof("Enquing Response %s", t1.Type())
	t2, err := tasks.NewReponseTask("+11234567890", "YOOO")
	if err != nil {
		logger.Fatal(err)
	}
	info, err = client.Enqueue(t2)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("Successfully enqueued task %s with ID %v ", info.Type, info.ID)

}
