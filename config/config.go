package config

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
	"github.com/isaacRevan24/gamification-toolkit-logic/handler"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

func RunConfigs() {
	loggingSetup()
}

func loggingSetup() {
	logs := utility.NewLogging()
	handler.Logs = logs
	controller.Logs = logs
	repository.Logs = logs
}
