package main

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/deyvidm/sms/cmd/dispatcher/client"
	"github.com/deyvidm/sms/cmd/dispatcher/log"
	"github.com/deyvidm/sms/cmd/dispatcher/workers"
	"github.com/deyvidm/sms/pkg/tasks"
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

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		logger.Fatalf("unable to load AWS SDK config, %v", err)
	}
	pp := pinpoint.NewFromConfig(cfg)
	irs := client.NewInviteResponseStore(redisClient)
	backendClient := client.NewWebBackendClient(os.Getenv("SECRET"))

	dispatcher := workers.NewMessageDispatcher(backendClient, irs, pp)
	reponseProcessor := workers.NewResponseProcessor(backendClient, irs)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeNewInvite, dispatcher.HandleSendInviteTask)
	mux.HandleFunc(tasks.TypeNewResponse, reponseProcessor.HandleResponse)
	mux.HandleFunc(tasks.TypeNewMessage, dispatcher.HandleNewMessageTask)

	if err := srv.Run(mux); err != nil {
		logger.Fatal(err)
	}
}
