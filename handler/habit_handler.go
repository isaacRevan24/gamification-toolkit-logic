package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

func addHabit(context *gin.Context) {
	const functionName string = "addHabit"
	Logs.LogDebug("Start " + functionName)

	var addNewHabitRequest model.AddNewHabitRequest
	err := utility.GenericRequestJsonMapper(&addNewHabitRequest, context)
	if err != nil {
		Logs.LogError(err)
		return
	}

	var controller controller.HabitControllerInterface = controller.NewHabitController()
	response := controller.AddNewHabitController(addNewHabitRequest)

	Logs.LogDebug("End " + functionName)
	context.JSON(getHttpStatusByCode(response.Status.Code), response)
}
