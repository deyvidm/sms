package log

import (
	"io"
	"log"
	"os"
	"strings"
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
		logger.AddHook(&indentLogHook{})
	})
	return logger
}

// Define the indentLogHook struct implementing logrus.Hook interface
type indentLogHook struct{}

// Fire is called before a log entry is written
func (hook *indentLogHook) Fire(entry *logrus.Entry) error {
	entry.Message = strings.ReplaceAll(entry.Message, "\t", "   ")
	return nil
}

func (hook *indentLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
