//go:generate go run github.com/golang/mock/mockgen -source router.go -destination mock/router_mock.go -package mock
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/handler"
)

var (
	userHandler  handler.UserHandlerInterface  = handler.NewUserHandler()
	habitHandler handler.HabitHandlerInterface = handler.NewHabitHandler()
)

type RouterRegistersInterface interface {
	HabitRegister(router *gin.RouterGroup)
	UserRegister(router *gin.RouterGroup)
}

type routerRegisters struct{}

func NewHandlerRegisters() RouterRegistersInterface {
	return &routerRegisters{}
}

func (*routerRegisters) UserRegister(router *gin.RouterGroup) {
	router.POST("/sign-up", userHandler.SignUp)
}

func (*routerRegisters) HabitRegister(router *gin.RouterGroup) {
	router.POST("/new", habitHandler.AddHabit)
	router.DELETE("/:id/user/:user", habitHandler.DeleteHabit)
}
