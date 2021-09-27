package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
)

func (*userController) SignUpController(request model.SignUpRequest) model.SignUpResponse {
	const functionName string = "SignUpController"
	Logs.LogDebug("Start " + functionName)

	repo, _ = repository.GetConnection()
	var signUpResponse model.SignUpResponse
	err := repo.SignUpRepository(request.ID)

	if err != nil {
		signUpResponse.Status.Message = "Error saving the user"
		signUpResponse.Status.Code = model.BAD_REQUEST_ERROR_STATUS

		Logs.LogError(err)
		Logs.LogDebug("End " + functionName)
		return signUpResponse
	}

	signUpResponse.Status.Message = "Successfully saved user."
	signUpResponse.Status.Code = model.SUCCESS_CODE_STATUS

	Logs.LogInfo("Successfully saved user.")
	Logs.LogDebug("End " + functionName)

	return signUpResponse
}
