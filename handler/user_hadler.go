package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/sign-up", signUp)
}

func signUp(context *gin.Context) {
	var signUpRequest model.SignUpModel
	utility.GenericRequestJsonMapper(&signUpRequest, context)
	var repo repository.Repo
	databaseResponse := repo.SignUp(signUpRequest.ID)
	if databaseResponse != nil {
		context.JSON(http.StatusOK, gin.H{"status": "Error guardando al usuario."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "User saves successfully."})
}
