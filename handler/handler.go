package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

var (
	Logs utility.LoggingInterface
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/sign-up", signUp)
}

func HabitRegister(router *gin.RouterGroup) {
	router.POST("/new", addHabit)
}

func getHttpStatusByCode(code string) int {
	switch code {
	case model.SUCCESS_CODE_STATUS:
		return http.StatusOK
	case model.BAD_REQUEST_ERROR_STATUS:
		return http.StatusBadRequest
	case model.INTERNAL_SERVER_ERROR_STATUS:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
