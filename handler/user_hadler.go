package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/repository"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/sign-up", saveUser)
}

func saveUser(c *gin.Context) {
	var repo repository.Repo
	repo.SignUp("testId11111112")
	c.JSON(http.StatusOK, gin.H{"status": "User saves successfully."})
}
