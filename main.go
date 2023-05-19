package main

import (
	"os"

	"github.com/deyvidm/sms-asynq/client"
	"github.com/deyvidm/sms-asynq/log"
	"github.com/deyvidm/sms-asynq/tasks"
	"github.com/deyvidm/sms-asynq/workers"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
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
	if err := godotenv.Load(".env"); err != nil {
		logger.Fatal(err)
	}
	backendClient := client.New(os.Getenv("SECRET"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // default DB 0 is for asynq
	})
	dispatcher := workers.NewMessageDispatcher(backendClient, rdb)
	reponseProcessor := workers.NewResponseProcessor(backendClient, rdb)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeNewInvite, dispatcher.HandleSendInviteTask)
	mux.HandleFunc(tasks.TypeNewResponse, reponseProcessor.HandleResponse)

	if err := srv.Run(mux); err != nil {
		logger.Fatal(err)
	}
}
