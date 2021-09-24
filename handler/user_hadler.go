package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/sign-up", saveUser)
}

func saveUser(c *gin.Context) {
	var signUpRequest model.SignUpModel
	utility.GenericRequestJsonMapper(&signUpRequest, c)
	fmt.Println(signUpRequest)
	var repo repository.Repo
	repo.SignUp(signUpRequest.ID)
	c.JSON(http.StatusOK, gin.H{"status": "User saves successfully."})
}
