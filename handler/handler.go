package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/sign-up", signUp)
	router.POST("/habit", addHabit)
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
