package module

import (
	"github.com/comfortablynumb/goginrestapi/internal/componentregistry"
	"github.com/comfortablynumb/goginrestapi/internal/errorhandler"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// Interfaces

type Module interface {
	GetName() string
	SetUpComponents(errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry)
	SetUpRouter(errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry, router *gin.Engine)
	SetUpValidator(errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry, validator *validator.Validate)
}
