package handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/controller"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
	"github.com/isaacRevan24/gamification-toolkit-logic/utility"
)

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
