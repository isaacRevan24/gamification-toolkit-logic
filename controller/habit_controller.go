package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
)

func (*habitController) AddNewHabitController(reques model.AddNewHabitRequest) model.AddNewHabitResponse {
	const functionName string = "AddNewHabitController"
	Logs.LogDebug("Start " + functionName)

	repo, _ = repository.GetConnection()
	var response model.AddNewHabitResponse
	habitId, err := repo.AddNewHabitRepository(reques.UserId, reques.Name, reques.Description, reques.Condition, reques.Repetition)
	if err != nil {

		response.Status.Code = model.BAD_REQUEST_ERROR_STATUS
		response.Status.Message = "Error creating new habit."

		Logs.LogDebug("End " + functionName)
		return response
	}
	response.Status.Code = model.SUCCESS_CODE_STATUS
	response.Status.Message = "New habit created"
	response.HabitId = habitId

	Logs.LogInfo("New habit created")
	Logs.LogDebug("End " + functionName)

	return response
}
