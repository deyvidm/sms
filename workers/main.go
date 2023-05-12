package main

import (
	"github.com/deyvidm/sms-asynq/log"
	"github.com/deyvidm/sms-asynq/task"
	"github.com/hibiken/asynq"
)

func main() {
	logger := log.GetLogger()
	logger.Info("Starting asynq worker server")
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 10,
			Logger:      logger,
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeNewMessage, task.HandleNewMessageTask)

	if err := srv.Run(mux); err != nil {
		logger.Fatal(err)
	}
}
