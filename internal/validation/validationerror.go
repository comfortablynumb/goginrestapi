package validation

import "fmt"

// Struct

type ValidationError struct {
	Field     string
	Validator string
	Message   string
}

func (v *ValidationError) String() string {
	return fmt.Sprintf("Field: %s - Validator: %s - Message: %s", v.Field, v.Validator, v.Message)
}

// Static functions

func NewValidationError(field string, validator string, message string) *ValidationError {
	return &ValidationError{
		Field:     field,
		Validator: validator,
		Message:   message,
	}
}
