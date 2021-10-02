package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
)

func (*habitHandler) AddHabit(context *gin.Context) {
	const functionName string = "addHabit"
	Logs.LogDebug("Start " + functionName)

	var addNewHabitRequest model.AddNewHabitRequest
	err := mapper.GenericRequestJsonMapper(&addNewHabitRequest, context)
	if err != nil {
		Logs.LogError(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Missing arguments."})
		return
	}

	var controller controller.HabitControllerInterface = controller.NewHabitController()
	response := controller.AddNewHabitController(addNewHabitRequest)

	Logs.LogDebug("End " + functionName)
	context.JSON(getHttpStatusByCode(response.Status.Code), response)
}
