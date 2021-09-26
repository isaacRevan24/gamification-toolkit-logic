package controller

import (
	"log"
	"os"

	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
)

func (*habitController) AddNewHabitController(reques model.AddNewHabitRequest) model.AddNewHabitResponse {

	repo, _ = repository.GetConnection()
	var response model.AddNewHabitResponse
	habitId, err := repo.AddNewHabit(reques.UserId, reques.Name, reques.Description, reques.Condition, reques.Repetition)
	if err != nil {
		var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
		Error.Println(err)
		response.Status.Code = model.BAD_REQUEST_ERROR_STATUS
		response.Status.Message = "Error creating new habit."

		return response
	}
	response.Status.Code = model.SUCCESS_CODE_STATUS
	response.Status.Message = "New habit created"
	response.HabitId = habitId

	return response
}
