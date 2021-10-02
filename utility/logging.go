//go:generate go run github.com/golang/mock/mockgen -source logging.go -destination mock/logging_mock.go -package mock
package utility

import (
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type LoggingInterface interface {
	LogError(err error)
	LogInfo(message string)
	LogDebug(message string)
	LogWarning(message interface{})
}

type Logging struct{}

func NewLogging() LoggingInterface {

	log.SetOutput(os.Stderr)
	log.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	return &Logging{}
}

func (*Logging) LogError(err error) {
	log.Error(err)
}

func (*Logging) LogInfo(message string) {
	log.Info(message)
}

func (*Logging) LogDebug(message string) {
	log.SetLevel(log.DebugLevel)
	log.Debug(message)
}

func (*Logging) LogWarning(message interface{}) {
	log.Warning(message)
}

// TODO: Add debug, info, warning, panic
