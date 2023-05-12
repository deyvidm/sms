package log

import (
	"io"
	"log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

func GetLogger() *logrus.Logger {
	once.Do(func() {
		f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Failed to initialize logger:", err)
		}
		multi := io.MultiWriter(f, os.Stdout)
		logger = logrus.New()
		logger.SetOutput(multi)

	})
	return logger
}
