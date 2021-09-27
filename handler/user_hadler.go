package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

func signUp(context *gin.Context) {
	const functionName string = "signUp"
	Logs.LogDebug("Start " + functionName)

	var signUpRequest model.SignUpRequest
	err := utility.GenericRequestJsonMapper(&signUpRequest, context)

	if err != nil {
		Logs.LogError(err)
		return
	}

	var controller controller.UserControllerInterface = controller.NewUserController()
	response := controller.SignUpController(signUpRequest)

	Logs.LogDebug("End " + functionName)
	context.JSON(getHttpStatusByCode(response.Status.Code), response)
}
