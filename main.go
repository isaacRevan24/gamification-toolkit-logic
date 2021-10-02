package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/config"
	"github.com/isaacRevan24/gamification-toolkit-logic/handler"
)

var (
	userHandler  handler.UserHandlerInterface  = handler.NewUserHandler()
	habitHandler handler.HabitHandlerInterface = handler.NewHabitHandler()
)

func main() {

	r := gin.Default()

	config.RunConfigs()

	v1 := r.Group("/v1")

	userHandler.UserRegister(v1.Group("/user"))
	habitHandler.HabitRegister(v1.Group("/habit"))

	r.Run()
}
