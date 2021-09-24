package utility

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Funcion generica que parsea el json del request a el puntero de un struct request.
func GenericRequestJsonMapper(request interface{}, context *gin.Context) {
	if err := context.ShouldBindJSON(request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Missing arguments."})
		return
	}
}
