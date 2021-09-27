package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/handler"
)

func main() {

	r := gin.Default()

	// TODO: sertup global logging

	v1 := r.Group("/v1")

	handler.UserRegister(v1.Group("/user"))
	handler.HabitRegister(v1.Group("/habit"))

	r.Run()
}
