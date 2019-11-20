package types

import (
	"net/http"
)

type StatusError struct {
	Status  int
	Message string
}

func (se StatusError) Error() string {
	return se.Message
}

func NewMissingEntityError(id string) error {
	return StatusError{
		Status:  http.StatusNotFound,
		Message: "no entity found for ID: " + id,
	}
}
