package handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

func signUp(context *gin.Context) {
	var signUpRequest model.SignUpRequest
	err := utility.GenericRequestJsonMapper(&signUpRequest, context)
	if err != nil {
		// TODO: hacer paquete de logs
		var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
		Error.Println(err)
		return
	}
	var controller controller.UserControllerInterface = controller.NewUserController()
	response := controller.SignUpController(signUpRequest)
	context.JSON(getHttpStatusByCode(response.Status.Code), response)
}
