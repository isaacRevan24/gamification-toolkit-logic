package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/config"
	"github.com/isaacRevan24/gamification-toolkit-logic/router"
)

var (
	gamificationRouter router.RouterRegistersInterface = router.NewHandlerRegisters()
)

func main() {

	r := gin.Default()

	config.RunConfigs()

	v1 := r.Group("/v1")

	gamificationRouter.UserRegister(v1.Group("/user"))
	gamificationRouter.HabitRegister(v1.Group("/habit"))

	r.Run()
}
