package main

import (
	"os"

	"github.com/deyvidm/sms/common/tasks"
	"github.com/deyvidm/sms/dispatcher/client"
	"github.com/deyvidm/sms/dispatcher/log"
	"github.com/deyvidm/sms/dispatcher/workers"
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

	// redis client for managing non-asynq data i.e. invite queuing
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // default DB 0 is for asynq
	})

	if err := godotenv.Load(".env"); err != nil {
		logger.Fatal(err)
	}

	irs := client.NewInviteResponseStore(redisClient)
	backendClient := client.NewWebBackendClient(os.Getenv("SECRET"))

	dispatcher := workers.NewMessageDispatcher(backendClient, irs)
	reponseProcessor := workers.NewResponseProcessor(backendClient, irs)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeNewInvite, dispatcher.HandleSendInviteTask)
	mux.HandleFunc(tasks.TypeNewResponse, reponseProcessor.HandleResponse)

	if err := srv.Run(mux); err != nil {
		logger.Fatal(err)
	}
}
