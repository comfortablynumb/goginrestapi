package middleware

import (
	"github.com/comfortablynumb/goginrestapi/internal/errorhandler"
	"github.com/gin-gonic/gin"
)

// Static functions

func ErrorHandler(errType gin.ErrorType, errorHandler *errorhandler.ErrorHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) < 1 {
			return
		}

		err := detectedErrors[0].Err

		controllerError := errorHandler.CreateHttpErrorFromErr(err)

		c.AbortWithStatusJSON(controllerError.HttpStatus, controllerError)

		return
	}
}
