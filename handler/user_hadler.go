package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

func UserRegister(router *gin.RouterGroup) {
	router.POST("/sign-up", signUp)
	router.POST("/habit", addHabit)
}

func signUp(context *gin.Context) {
	var signUpRequest model.SignUpRequest
	mapping := utility.GenericRequestJsonMapper(&signUpRequest, context)
	if mapping != nil {
		// TODO: hacer paquete de logs
		var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
		Error.Println(mapping)
		return
	}
	var controller controller.UserControllerInterface = controller.NewUserController()
	response := controller.SignUpController(signUpRequest)
	context.JSON(getHttpStatusByCode(response.Status.Code), response)
}

func addHabit(context *gin.Context) {
	var addNewHabitRequest model.AddNewHabitRequest
	mapping := utility.GenericRequestJsonMapper(&addNewHabitRequest, context)
	if mapping != nil {
		// TODO: hacer paquete de logs
		var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
		Error.Println(mapping)

		return
	}
	var controller controller.UserControllerInterface = controller.NewUserController()
	response := controller.AddNewHabitController(addNewHabitRequest)
	context.JSON(getHttpStatusByCode(response.Status.Code), response)
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
