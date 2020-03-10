package validation

// Struct

type ValidationError struct {
	Field     string
	Validator string
	Message   string
}

// Static functions

func NewValidationError(field string, validator string, message string) *ValidationError {
	return &ValidationError{
		Field:     field,
		Validator: validator,
		Message:   message,
	}
}
