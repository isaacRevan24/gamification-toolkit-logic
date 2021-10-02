//go:generate go run github.com/golang/mock/mockgen -source mapper.go -destination mock/mapper_mock.go -package mock
package utility

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type GamificationMapper interface {
	GenericRequestJsonMapper(request interface{}, context *gin.Context) error
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
