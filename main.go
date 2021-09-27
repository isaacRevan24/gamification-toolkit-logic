package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/config"
	"github.com/isaacRevan24/gamification-toolkit-logic/handler"
)

func main() {

	r := gin.Default()

	config.RunConfigs()

	v1 := r.Group("/v1")

	handler.UserRegister(v1.Group("/user"))
	handler.HabitRegister(v1.Group("/habit"))

	r.Run()
}
