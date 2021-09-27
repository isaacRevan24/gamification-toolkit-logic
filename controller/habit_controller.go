package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
)

func (*habitController) AddNewHabitController(reques model.AddNewHabitRequest) model.AddNewHabitResponse {

	repo, _ = repository.GetConnection()
	var response model.AddNewHabitResponse
	habitId, err := repo.AddNewHabit(reques.UserId, reques.Name, reques.Description, reques.Condition, reques.Repetition)
	if err != nil {

		Logs.LogError(err)

		response.Status.Code = model.BAD_REQUEST_ERROR_STATUS
		response.Status.Message = "Error creating new habit."

		return response
	}
	response.Status.Code = model.SUCCESS_CODE_STATUS
	response.Status.Message = "New habit created"
	response.HabitId = habitId
	Logs.LogInfo("New habit created")

	return response
}
