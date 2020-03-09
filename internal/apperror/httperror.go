package apperror

import (
	"fmt"
	"net/http"
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

func NewBindingHttpError(err error, source string) *HttpError {
	return NewHttpError(err, source, http.StatusBadRequest, BindingErrorCode, fmt.Sprintf("%s: %s", BindingErrorMessage, err.Error()), nil)
}

func NewValidationHttpError(err error, source string) *HttpError {
	return NewHttpError(err, source, http.StatusBadRequest, ValidationErrorCode, fmt.Sprintf("%s: %s", ValidationErrorMessage, err.Error()), nil)
}

func NewInternalServerHttpError(err error, source string) *HttpError {
	return NewHttpError(err, source, http.StatusInternalServerError, InternalErrorCode, InternalErrorMessage, nil)
}

func NewDbHttpError(err error, source string) *HttpError {
	return NewHttpError(err, source, http.StatusInternalServerError, DbErrorCode, DbErrorMessage, nil)
}

func NewNotFoundHttpError(err error, source string) *HttpError {
	return NewHttpError(err, source, http.StatusNotFound, ModelNotFoundErrorCode, ModelNotFoundErrorMessage, nil)
}

func NewHttpError(err error, source string, httpStatus int, code string, message string, data map[string]interface{}) *HttpError {
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
