package utility

import (
	log "github.com/sirupsen/logrus"
)

type LoggingInterface interface {
	LogError(err error)
	LogInfo(message string)
	LogDebug(message interface{})
	LogWarning(message interface{})
}

type Logging struct{}

func NewLogging() LoggingInterface {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	return &Logging{}
}

func (*Logging) LogError(err error) {
	log.Error(err)
}

func (*Logging) LogInfo(message string) {
	log.Info(message)
}

func (*Logging) LogDebug(message interface{}) {
	log.Debug(message)
}

func (*Logging) LogWarning(message interface{}) {
	log.Warning(message)
}

// TODO: Add debug, info, warning, panic
