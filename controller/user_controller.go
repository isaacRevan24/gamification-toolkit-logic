package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
)

type UserController struct {
}

func (controller UserController) SignUpController(request model.SignUpRequest) model.SignUpResponse {

	var repo repository.Repo
	var signUpResponse model.SignUpResponse
	databaseResponse := repo.SignUp(request.ID)

	if databaseResponse != nil {
		signUpResponse.Status = "Error saving the user"
		signUpResponse.Code = model.BAD_REQUEST_ERROR_STATUS
		return signUpResponse
	}

	signUpResponse.Status = "Successfully saved user."
	signUpResponse.Code = model.SUCCESS_CODE_STATUS
	return signUpResponse
}
