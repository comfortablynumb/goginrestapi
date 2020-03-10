package apperror

import (
	"fmt"
	"net/http"

	"github.com/comfortablynumb/goginrestapi/internal/context"
)

// Structs

type HttpError struct {
	Err        error                  `json:"-"`
	HttpStatus int                    `json:"-"`
	Source     string                 `json:"-"`
	Code       string                 `json:"code"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("[%s] Http Status: %d - Code: %s - Message: %s", e.Source, e.HttpStatus, e.Code, e.Message)
}

func (e *HttpError) String() string {
	return e.Error()
}

// Static functions

func NewBindingHttpError(ctx *context.RequestContext, err error, source string, data map[string]interface{}) *HttpError {
	if data == nil {
		data = make(map[string]interface{})
	}

	AddValidationErrorsToMap(ctx, err, data)

	return NewHttpError(ctx, err, source, http.StatusBadRequest, BindingErrorCode, BindingErrorMessage, data)
}

func NewValidationHttpError(ctx *context.RequestContext, err error, source string, data map[string]interface{}) *HttpError {
	return NewHttpError(ctx, err, source, http.StatusBadRequest, ValidationErrorCode, ValidationErrorMessage, data)
}

func NewInternalServerHttpError(ctx *context.RequestContext, err error, source string, data map[string]interface{}) *HttpError {
	return NewHttpError(ctx, err, source, http.StatusInternalServerError, InternalErrorCode, InternalErrorMessage, data)
}

func NewDbHttpError(ctx *context.RequestContext, err error, source string, data map[string]interface{}) *HttpError {
	return NewHttpError(ctx, err, source, http.StatusInternalServerError, DbErrorCode, DbErrorMessage, data)
}

func NewNotFoundHttpError(ctx *context.RequestContext, err error, source string, data map[string]interface{}) *HttpError {
	return NewHttpError(ctx, err, source, http.StatusNotFound, ModelNotFoundErrorCode, ModelNotFoundErrorMessage, data)
}

func NewHttpError(ctx *context.RequestContext, err error, source string, httpStatus int, code string, message string, data map[string]interface{}) *HttpError {
	if data == nil {
		data = make(map[string]interface{})
	}

	return &HttpError{
		Err:        err,
		HttpStatus: httpStatus,
		Source:     source,
		Code:       code,
		Message:    message,
		Data:       data,
	}
}
