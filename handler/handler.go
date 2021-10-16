//go:generate go run github.com/golang/mock/mockgen -source handler.go -destination mock/handler_mock.go -package mock
package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

var (
	mapper          utility.GamificationMapper = utility.NewGamificationMapper()
	Logs            utility.LoggingInterface
	userRepository  repository.UserRepository           = repository.NewUserRepository()
	habitRepository repository.HabitRepository          = repository.NewHabitRepository()
	habitController controller.HabitControllerInterface = controller.NewHabitController(habitRepository)
	userController  controller.UserControllerInterface  = controller.NewUserController(userRepository)
)

type UserHandlerInterface interface {
	SignUp(context *gin.Context)
}

type HabitHandlerInterface interface {
	AddHabit(context *gin.Context)
	DeleteHabit(context *gin.Context)
	CheckHabit(context *gin.Context)
}

type userHandler struct{}

type habitHandler struct{}

func NewHabitHandler() HabitHandlerInterface {
	return &habitHandler{}
}

func NewUserHandler() UserHandlerInterface {
	return &userHandler{}
}

func (*userHandler) SignUp(context *gin.Context) {
	const functionName string = "signUp"
	Logs.LogDebug("Start " + functionName)

	var signUpRequest model.SignUpRequest
	var response model.SignUpResponse
	mapperError := mapper.GenericRequestJsonMapper(&signUpRequest, context)

	if mapperError != nil {
		Logs.LogError(mapperError)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Missing arguments."})
		return
	}

	signUpError := userController.SignUp(signUpRequest)

	if signUpError != nil {
		response.Status = mapper.StatusBuilder(model.BAD_REQUEST_ERROR_STATUS, "Error saving the user")
		context.JSON(http.StatusBadRequest, response)

		Logs.LogError(signUpError)
		Logs.LogDebug("End " + functionName)
		return
	}

	response.Status = mapper.StatusBuilder(model.SUCCESS_CODE_STATUS, "Successfully saved user.")

	Logs.LogDebug("End " + functionName)
	context.JSON(http.StatusOK, response)
}

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

func (*habitHandler) CheckHabit(context *gin.Context) {
	var headers model.CheckHabitHeaders

	if err := context.ShouldBindHeader(&headers); err != nil {
		context.JSON(200, err)
	}

	fmt.Println(headers)

	context.JSON(http.StatusOK, gin.H{"chech habit": "ok"})
}
