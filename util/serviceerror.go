package util

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type ServiceError struct {
	Id      string
	Message string
	Sub     error
	Code    int
}

type NotFoundError struct {
	error
}

func NewServiceError(message string, code int, sub error) *ServiceError {
	return &ServiceError{
		Id:      uuid.New().String(),
		Message: message,
		Sub:     sub,
		Code:    code,
	}
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("ERROR[%s:%s]%s", e.Id, e.Message, e.Sub.Error())
}

func (e *ServiceError) UserError() string {
	return fmt.Sprintf("%s:%s", e.Id, e.Message)
}

/*
	Maneja el Error.

	Si el error es nil entonces no hace nada y devuelve falso
	Si el error es un ServiceError entonces usa el codigo que tiene
	Si el error no es un Service Error lo considera como un SERVER_ERROR con un status de InternalServerError
*/
func LogError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}
	sErr, ok := err.(*ServiceError)
	if !ok {
		sErr = NewServiceError("SERVER_ERROR", http.StatusInternalServerError, err)
	}

	log.Println(sErr)
	http.Error(w, sErr.UserError(), sErr.Code)
	return true
}
