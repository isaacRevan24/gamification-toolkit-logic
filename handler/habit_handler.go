package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
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

	if controllerError == controller.ErrorInvalidHabitCondition {

		response.Status = mapper.StatusBuilder(model.BAD_REQUEST_ERROR_STATUS, "Invalid habit condition.")

		Logs.LogError(controllerError)
		context.JSON(http.StatusBadRequest, response)
		return

	} else if controllerError != nil {

		response.Status = mapper.StatusBuilder(model.BAD_REQUEST_ERROR_STATUS, "Error creating new habit.")

		Logs.LogError(controllerError)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response.Status = mapper.StatusBuilder(model.SUCCESS_CODE_STATUS, "New habit created")
	response.HabitId = habitId

	Logs.LogDebug("End " + functionName)

	context.JSON(http.StatusOK, response)
}

func (*habitHandler) DeleteHabit(context *gin.Context) {
	const functionName string = "DeleteHabit"
	Logs.LogDebug("Start " + functionName)

	habitId := context.Param("id")
	userId := context.Param("user")
	Logs.LogDebug("Habit id: " + habitId)
	Logs.LogDebug("user id: " + userId)

	var response model.DeleteHabitResponse

	isDeleted, controllerError := habitRepository.DeleteHabitRepository(userId, habitId)

	if controllerError != nil {
		response.Status = mapper.StatusBuilder(model.BAD_REQUEST_ERROR_STATUS, "Error deleting habit.")
		Logs.LogError(controllerError)
		Logs.LogDebug("End " + functionName)
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	if isDeleted {
		response.Status = mapper.StatusBuilder(model.SUCCESS_CODE_STATUS, "Habit deleted.")
		Logs.LogDebug("End " + functionName)
		context.JSON(http.StatusOK, response)
		return
	} else {
		response.Status = mapper.StatusBuilder(model.BAD_REQUEST_ERROR_STATUS, "Habit already deleted.")
		Logs.LogDebug("End " + functionName)
		context.JSON(http.StatusBadRequest, response)
		return
	}

}
