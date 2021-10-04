//go:generate go run github.com/golang/mock/mockgen -source handler.go -destination mock/handler_mock.go -package mock
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

var (
	mapper          utility.GamificationMapper = utility.NewGamificationMapper()
	Logs            utility.LoggingInterface
	userRepository  repository.UserRepository           = repository.NewUserRepository()
	habitRepository repository.HabitRepository          = repository.NewHabitRepository()
	habitController controller.HabitControllerInterface = controller.NewHabitController(habitRepository)
	userController  controller.UserControllerInterface  = controller.NewUserController(userRepository)
)

type UserHandlerInterface interface {
	SignUp(context *gin.Context)
}

type HabitHandlerInterface interface {
	AddHabit(context *gin.Context)
	DeleteHabit(context *gin.Context)
}

type userHandler struct{}

type habitHandler struct{}

func NewHabitHandler() HabitHandlerInterface {
	return &habitHandler{}
}
func NewUserHandler() UserHandlerInterface {
	return &userHandler{}
}
