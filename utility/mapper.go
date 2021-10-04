//go:generate go run github.com/golang/mock/mockgen -source mapper.go -destination mock/mapper_mock.go -package mock
package utility

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
)

type GamificationMapper interface {
	GenericRequestJsonMapper(request interface{}, context *gin.Context) error
	StatusBuilder(code string, message string) model.StatusResponse
}

type mapper struct{}

func NewGamificationMapper() GamificationMapper {
	return &mapper{}
}

// Funcion generica que parsea el json del request a el puntero de un struct request.
func (*mapper) GenericRequestJsonMapper(request interface{}, context *gin.Context) error {
	if err := context.ShouldBindJSON(request); err != nil {
		return errors.New("missing argument")
	}
	return nil
}

func (*mapper) StatusBuilder(code string, message string) model.StatusResponse {
	status := model.StatusResponse{Code: code, Message: message}
	return status
}
