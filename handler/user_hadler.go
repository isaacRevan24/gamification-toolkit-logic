package handler

import (
	"net/http"

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
		context.JSON(http.StatusBadRequest, gin.H{"error": "Missing arguments."})
		return
	}

	var controller controller.UserControllerInterface = controller.NewUserController()
	response := controller.SignUpController(signUpRequest)

	Logs.LogDebug("End " + functionName)
	context.JSON(getHttpStatusByCode(response.Status.Code), response)
}
