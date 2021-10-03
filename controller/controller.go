//go:generate go run github.com/golang/mock/mockgen -source controller.go -destination mock/controller_mock.go -package mock
package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

var (
	userRepository  repository.UserRepository
	habitRepository repository.HandlerRepository
	Logs            utility.LoggingInterface = utility.NewLogging()
)

type UserControllerInterface interface {
	SignUpController(request model.SignUpRequest) model.SignUpResponse
}
type HabitControllerInterface interface {
	AddNewHabitController(reques model.AddNewHabitRequest) model.AddNewHabitResponse
}

type userController struct{}

type habitController struct{}

func NewUserController(repo repository.UserRepository) UserControllerInterface {
	userRepository = repo
	return &userController{}
}

func NewHabitController(repo repository.HandlerRepository) HabitControllerInterface {
	habitRepository = repo
	return &habitController{}
}
