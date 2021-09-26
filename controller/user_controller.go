package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
)

type UserControllerInterface interface {
	SignUpController(request model.SignUpRequest) model.SignUpResponse
}

type userController struct{}

var (
	repo *repository.Repo
)

func NewUserController() UserControllerInterface {
	return &userController{}
}

func (*userController) SignUpController(request model.SignUpRequest) model.SignUpResponse {

	repo, _ = repository.GetConnection()
	var signUpResponse model.SignUpResponse
	databaseResponse := repo.SignUp(request.ID)

	if databaseResponse != nil {
		signUpResponse.Message = "Error saving the user"
		signUpResponse.Code = model.BAD_REQUEST_ERROR_STATUS
		return signUpResponse
	}

	signUpResponse.Message = "Successfully saved user."
	signUpResponse.Code = model.SUCCESS_CODE_STATUS
	return signUpResponse
}
