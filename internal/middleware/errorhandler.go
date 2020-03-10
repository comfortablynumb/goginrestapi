package middleware

import (
	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/errorhandler"
	"github.com/gin-gonic/gin"
)

// Static functions

func ErrorHandler(
	requestContextFactory *context.RequestContextFactory,
	errType gin.ErrorType,
	errorHandler *errorhandler.ErrorHandler,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) < 1 {
			return
		}

		err := detectedErrors[0].Err

		languages := c.GetHeader("Accept-Language")

		controllerError := errorHandler.CreateHttpErrorFromErr(requestContextFactory.NewRequestContext(c), err, languages)

		c.AbortWithStatusJSON(controllerError.HttpStatus, controllerError)

		return
	}
}
