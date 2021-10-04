package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
)

func (*habitHandler) AddHabit(context *gin.Context) {
	const functionName string = "AddHabit"
	Logs.LogDebug("Start " + functionName)

	var addNewHabitRequest model.AddNewHabitRequest
	var response model.AddNewHabitResponse

	mappingError := mapper.GenericRequestJsonMapper(&addNewHabitRequest, context)

	if mappingError != nil {
		Logs.LogError(mappingError)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Missing arguments."})
		return
	}

	habitId, controllerError := habitController.AddNewHabit(addNewHabitRequest)

	if controllerError != nil {
		response.Status.Code = model.BAD_REQUEST_ERROR_STATUS
		response.Status.Message = "Error creating new habit."
		Logs.LogError(controllerError)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response.Status.Code = model.SUCCESS_CODE_STATUS
	response.Status.Message = "New habit created"
	response.HabitId = habitId

	Logs.LogDebug("End " + functionName)

	context.JSON(http.StatusOK, response)
}

func (*habitHandler) DeleteHabit(context *gin.Context) {
	const functionName string = "DeleteHabit"
	Logs.LogDebug("Start " + functionName)

	var deleteHabitRequest model.DeleteHabitRequest
	//var deleteHabitResponse model.DeleteHabitResponse

	mappingError := mapper.GenericRequestJsonMapper(&deleteHabitRequest, context)

	if mappingError != nil {
		Logs.LogError(mappingError)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Missing arguments."})
		return
	}
	Logs.LogDebug("End " + functionName)
	context.JSON(http.StatusOK, gin.H{"Status": "Habit deleted"})
}
