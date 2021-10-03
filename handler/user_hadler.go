package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
)

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
		response.Status.Code = model.BAD_REQUEST_ERROR_STATUS
		response.Status.Message = "Error saving the user"
		context.JSON(http.StatusBadRequest, response)
		return
	}

	response.Status.Code = model.SUCCESS_CODE_STATUS
	response.Status.Message = "Successfully saved user."

	Logs.LogDebug("End " + functionName)
	context.JSON(http.StatusOK, response)
}
