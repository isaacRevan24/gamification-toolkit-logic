package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
)

var (
	repo *repository.Repo
)

type UserControllerInterface interface {
	SignUpController(request model.SignUpRequest) model.SignUpResponse
}
type HabitControllerInterface interface {
	AddNewHabitController(reques model.AddNewHabitRequest) model.AddNewHabitResponse
}

type userController struct{}

type habitController struct{}

func NewUserController() UserControllerInterface {
	return &userController{}
}

func NewHabitController() HabitControllerInterface {
	return &habitController{}
}
