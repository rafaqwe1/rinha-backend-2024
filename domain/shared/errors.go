package shared

import "errors"

var NotFoundError = errors.New("not found")
var NoRowsAffectedError = errors.New("")

var TypeValidationError *ValidationError

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}
