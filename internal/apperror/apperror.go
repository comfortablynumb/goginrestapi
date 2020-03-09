package apperror

import "fmt"

// Structs

type AppError struct {
	Err     error
	Source  string                 `json:"-"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewValidationAppError(err error, source string) *AppError {
	return NewAppError(err, source, ValidationErrorCode, ValidationErrorMessage, nil)
}

func NewDbAppError(err error, source string) *AppError {
	return NewAppError(err, source, DbErrorCode, err.Error(), nil)
}

func NewModelNotFoundAppError(err error, source string) *AppError {
	return NewAppError(err, source, ModelNotFoundErrorCode, ModelNotFoundErrorMessage, nil)
}

func NewAppError(err error, source string, code string, message string, data map[string]interface{}) *AppError {
	if data == nil {
		data = make(map[string]interface{})
	}

	return &AppError{
		Err:     err,
		Source:  source,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] Code: %s - Message: %s - Error: %s", e.Source, e.Code, e.Message, e.Err)
}

func (e *AppError) String() string {
	return e.Error()
}
