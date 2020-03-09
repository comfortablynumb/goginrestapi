package errorhandler

import (
	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	hooks2 "github.com/comfortablynumb/goginrestapi/internal/hooks"
	"github.com/rs/zerolog"
)

// Struct

type ErrorHandler struct {
	hooks  *hooks2.Hooks
	logger *zerolog.Logger
}

func (e *ErrorHandler) HandleFatal(err error, message string) {
	e.logger.Error().Msgf("[ERROR] %s - Error: %s", message, err)

	panic(err)
}

func (e *ErrorHandler) HandleFatalIfError(err error, message string) {
	if err == nil {
		return
	}

	e.HandleFatal(err, message)
}

func (e *ErrorHandler) MapAppErrorToHttpError(err *apperror.AppError) *apperror.HttpError {
	switch err.Code {
	case apperror.BindingErrorCode:
		return apperror.NewBindingHttpError(err.Err, err.Source)
	case apperror.ValidationErrorCode:
		return apperror.NewValidationHttpError(err.Err, err.Source)
	case apperror.DbErrorCode:
		return apperror.NewDbHttpError(err.Err, err.Source)
	case apperror.ModelNotFoundErrorCode:
		return apperror.NewNotFoundHttpError(err.Err, err.Source)
	default:
		return apperror.NewInternalServerHttpError(err.Err, err.Source)
	}
}

func (e *ErrorHandler) CreateHttpErrorFromErr(err error) *apperror.HttpError {
	var controllerError *apperror.HttpError

	switch err.(type) {
	case *apperror.AppError:
		appError := err.(*apperror.AppError)

		e.logger.Error().Msgf("[Application Error] %s", appError.String())

		controllerError = e.MapAppErrorToHttpError(err.(*apperror.AppError))
	case *apperror.HttpError:
		controllerError = err.(*apperror.HttpError)
	default:
		controllerError = apperror.NewInternalServerHttpError(err, "ErrorHandler")
	}

	return controllerError
}

// Static functions

func NewErrorHandler(logger *zerolog.Logger, hooks *hooks2.Hooks) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
		hooks:  hooks,
	}
}
