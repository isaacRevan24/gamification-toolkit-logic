//go:generate go run github.com/golang/mock/mockgen -source controller.go -destination mock/controller_mock.go -package mock
package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

var (
	userRepository  repository.UserRepository
	habitRepository repository.HabitRepository
	Logs            utility.LoggingInterface = utility.NewLogging()

	conditionTypes = [3]string{"D", "W", "Y"}
)

type UserControllerInterface interface {
	SignUp(request model.SignUpRequest) error
}
type HabitControllerInterface interface {
	AddNewHabit(reques model.AddNewHabitRequest) (int, error)
}

type userController struct{}

type habitController struct{}

func NewUserController(repo repository.UserRepository) UserControllerInterface {
	userRepository = repo
	return &userController{}
}

func NewHabitController(repo repository.HabitRepository) HabitControllerInterface {
	habitRepository = repo
	return &habitController{}
}
