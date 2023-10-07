package controllers

import (
	"sync"

	"github.com/hibiken/asynq"
)

var (
	asynqClient *asynq.Client
	once        sync.Once
)

func GetAsynqClient() *asynq.Client {
	once.Do(func() {
		asynqClient = asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	})
	return asynqClient
}
