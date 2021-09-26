package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
)

func (*habitController) AddNewHabitController(reques model.AddNewHabitRequest) model.AddNewHabitResponse {

	repo, _ = repository.GetConnection()
	var response model.AddNewHabitResponse
	repo.AddNewHabit("testUserid1", "new", "a", "Y", 2)
	response.Status.Code = model.SUCCESS_CODE_STATUS
	response.Status.Message = "New habit created"

	return response
}
