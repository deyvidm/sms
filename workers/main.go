package main

import (
	"os"

	"github.com/deyvidm/sms-asynq/client"
	"github.com/deyvidm/sms-asynq/log"
	"github.com/deyvidm/sms-asynq/task"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
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
	if err := godotenv.Load("../.env"); err != nil {
		logger.Fatal(err)
	}
	backendClient := client.New(os.Getenv("SECRET"))
	dispatcher := task.NewMessageDispatcher(backendClient)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeNewMessage, dispatcher.HandleNewMessageTask)

	if err := srv.Run(mux); err != nil {
		logger.Fatal(err)
	}
}
