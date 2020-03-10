package errorhandler

import (
	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/comfortablynumb/goginrestapi/internal/context"
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

func (e *ErrorHandler) MapAppErrorToHttpError(ctx *context.RequestContext, err *apperror.AppError) *apperror.HttpError {
	switch err.Code {
	case apperror.BindingErrorCode:
		return apperror.NewBindingHttpError(ctx, err.Err, err.Source)
	case apperror.ValidationErrorCode:
		return apperror.NewValidationHttpError(ctx, err.Err, err.Source)
	case apperror.DbErrorCode:
		return apperror.NewDbHttpError(ctx, err.Err, err.Source)
	case apperror.ModelNotFoundErrorCode:
		return apperror.NewNotFoundHttpError(ctx, err.Err, err.Source)
	default:
		return apperror.NewInternalServerHttpError(ctx, err.Err, err.Source)
	}
}

func (e *ErrorHandler) CreateHttpErrorFromErr(ctx *context.RequestContext, err error, MapAppErrorToHttpError string) *apperror.HttpError {
	var controllerError *apperror.HttpError

	switch err.(type) {
	case *apperror.AppError:
		appError := err.(*apperror.AppError)

		e.logger.Error().Msgf("[Application Error] %s", appError.String())

		controllerError = e.MapAppErrorToHttpError(ctx, err.(*apperror.AppError))
	case *apperror.HttpError:
		controllerError = err.(*apperror.HttpError)
	default:
		controllerError = apperror.NewInternalServerHttpError(ctx, err, "ErrorHandler")
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
