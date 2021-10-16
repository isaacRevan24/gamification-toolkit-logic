//go:generate go run github.com/golang/mock/mockgen -source controller.go -destination mock/controller_mock.go -package mock
package controller

import (
	"time"

	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

var (
	userRepository  repository.UserRepository
	habitRepository repository.HabitRepository
	Logs            utility.LoggingInterface = utility.NewLogging()
)

type UserControllerInterface interface {
	SignUp(request model.SignUpRequest) error
}
type HabitControllerInterface interface {
	CheckHabit(userId string, habitId int, date time.Time) (string, error)
	AddNewHabit(reques model.AddNewHabitRequest) (int, error)
	DeleteHabit(userId string, habitId string) (bool, error)
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

func (*userController) SignUp(request model.SignUpRequest) error {
	const functionName string = "SignUpController"
	Logs.LogDebug("Start " + functionName)

	err := userRepository.SignUpRepository(request.ID)

	if err != nil {
		Logs.LogDebug("End " + functionName)
		return err
	}

	Logs.LogInfo("Successfully saved user.")
	Logs.LogDebug("End " + functionName)

	return nil
}

func (*habitController) AddNewHabit(reques model.AddNewHabitRequest) (int, error) {
	const functionName string = "AddNewHabitController"
	Logs.LogDebug("Start " + functionName)

	conditionError := validateHabitCondition(reques.Condition)

	if conditionError != nil {
		Logs.LogError(conditionError)
		Logs.LogDebug("End " + functionName)
		return 0, conditionError
	}

	habitId, err := habitRepository.AddNewHabitRepository(reques.UserId, reques.Name, reques.Description, reques.Condition, reques.Repetition)

	if err != nil {
		Logs.LogDebug("End " + functionName)
		return 0, err
	}

	Logs.LogInfo("New habit created")
	Logs.LogDebug("End " + functionName)

	return habitId, nil
}

func (*habitController) DeleteHabit(userId string, habitId string) (bool, error) {
	const functionName string = "DeleteHabit"
	Logs.LogDebug("Start " + functionName)

	response, err := habitRepository.DeleteHabitRepository(userId, habitId)

	if err != nil {
		Logs.LogDebug("End " + functionName)
		return false, err
	}

	Logs.LogDebug("End " + functionName)
	return response, nil
}

func (*habitController) CheckHabit(userId string, habitId int, date time.Time) (string, error) {
	const functionName string = "DeleteHabit"
	Logs.LogDebug("Start " + functionName)
	Logs.LogDebug("End " + functionName)
	return "flaskdjfljfd", nil
}

func validateHabitCondition(condition string) error {

	conditionTypes := [3]string{"D", "W", "Y"}

	for _, i := range conditionTypes {
		if condition == i {
			return nil
		}
	}

	return ErrorInvalidHabitCondition
}
