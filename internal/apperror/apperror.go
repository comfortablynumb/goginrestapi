package apperror

import (
	"fmt"

	"github.com/comfortablynumb/goginrestapi/internal/context"
)

// Structs

type AppError struct {
	Err     error
	Source  string                 `json:"-"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewValidationAppError(ctx *context.RequestContext, err error, source string) *AppError {
	data := make(map[string]interface{})

	AddValidationErrorsToMap(ctx, err, data)

	return NewAppError(ctx, err, source, ValidationErrorCode, ValidationErrorMessage, data)
}

func NewDbAppError(ctx *context.RequestContext, err error, source string) *AppError {
	return NewAppError(ctx, err, source, DbErrorCode, err.Error(), nil)
}

func NewModelNotFoundAppError(ctx *context.RequestContext, err error, source string) *AppError {
	return NewAppError(ctx, err, source, ModelNotFoundErrorCode, ModelNotFoundErrorMessage, nil)
}

func NewAppError(ctx *context.RequestContext, err error, source string, code string, message string, data map[string]interface{}) *AppError {
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
