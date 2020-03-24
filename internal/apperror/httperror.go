package apperror

import (
	"fmt"
	"net/http"

	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/validation"
	"github.com/mitchellh/mapstructure"
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
	return fmt.Sprintf("[%s] Http Status: %d - Code: %s - Message: %s - Data: %v", e.Source, e.HttpStatus, e.Code, e.Message, e.Data)
}

func (e *HttpError) String() string {
	return e.Error()
}

func (e *HttpError) HasErrorCount(count int) bool {
	return e.GetErrorCount() == count
}

func (e *HttpError) GetErrors() []*validation.ValidationError {
	res := make([]*validation.ValidationError, 0)
	errors, found := e.Data["errors"]

	if !found {
		return res
	}

	errorArray, ok := errors.([]interface{})

	if !ok {
		return res
	}

	for _, element := range errorArray {
		validationError := &validation.ValidationError{}

		err := mapstructure.Decode(element, validationError)

		if err != nil {
			return res
		}

		res = append(res, validationError)
	}

	return res
}

func (e *HttpError) GetErrorCount() int {
	return len(e.GetErrors())
}

func (e *HttpError) HasErrorCountByNameAndType(count int, fieldName string, errorType string) bool {
	return e.GetErrorCountByNameAndType(fieldName, errorType) == count
}

func (e *HttpError) GetErrorCountByNameAndType(fieldName string, errorType string) int {
	count := 0

	for _, validationError := range e.GetErrors() {
		if validationError.Field == fieldName && validationError.Validator == errorType {
			count++
		}
	}

	return count
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
