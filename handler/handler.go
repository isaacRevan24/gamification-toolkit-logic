//go:generate go run github.com/golang/mock/mockgen -source handler.go -destination mock/handler_mock.go -package mock
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

var (
	mapper utility.GamificationMapper = utility.NewGamificationMapper()
	Logs   utility.LoggingInterface
)

type UserHandlerInterface interface {
	UserRegister(router *gin.RouterGroup)
}

type HabitHandlerInterface interface {
	HabitRegister(router *gin.RouterGroup)
}

type userHandler struct{}

type habitHandler struct{}

func NewUserHandler() UserHandlerInterface {
	return &userHandler{}
}

func NewHabitHandler() HabitHandlerInterface {
	return &habitHandler{}
}

func (*userHandler) UserRegister(router *gin.RouterGroup) {
	router.POST("/sign-up", signUp)
}

func (*habitHandler) HabitRegister(router *gin.RouterGroup) {
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
