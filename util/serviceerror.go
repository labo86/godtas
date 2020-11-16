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
}

func NewServiceError(message string, sub error) *ServiceError {
	return &ServiceError{
		Id:      uuid.New().String(),
		Message: message,
		Sub:     sub,
	}
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("ERROR[%s:%s]%s", e.Id, e.Message, e.Sub.Error())
}

func (e *ServiceError) UserError() string {
	return fmt.Sprintf("%s:%s", e.Id, e.Message)
}

func LogError(w http.ResponseWriter, err error) {
	sErr, ok := err.(*ServiceError)
	if !ok {
		sErr = NewServiceError("SERVER_ERROR", err)
	}

	log.Println(sErr)
	http.Error(w, sErr.UserError(), http.StatusInternalServerError)
}
