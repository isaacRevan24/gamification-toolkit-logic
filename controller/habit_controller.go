package controller

import "github.com/isaacRevan24/gamification-toolkit-logic/model"

func (*userController) AddNewHabitController(reques model.AddNewHabitRequest) model.AddNewHabitResponse {

	var response model.AddNewHabitResponse
	response.Status.Code = model.SUCCESS_CODE_STATUS
	response.Status.Message = "New habit created"

	return response
}
