package hooks

import (
	"github.com/comfortablynumb/goginrestapi/internal/componentregistry"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gopkg.in/go-playground/validator.v9"
)

// SetupRouter Method to customize the Gin router, or to return a completely new one.
func (h *Hooks) SetupRouter(router *gin.Engine) *gin.Engine {
	return router
}

// SetupComponentRegistry Method to customize the component registry, or to return a completely new one.
func (h *Hooks) SetupComponentRegistry(componentRegistry *componentregistry.ComponentRegistry) *componentregistry.ComponentRegistry {
	return componentRegistry
}

// SetupValidator Method to customize the validator, or to return a completely new one.
func (h *Hooks) SetupValidator(v *validator.Validate) *validator.Validate {
	return v
}

// SetupLogger Method to customize the logger, or to return a completely new one.
func (h *Hooks) SetupLogger(logger *zerolog.Logger) *zerolog.Logger {
	return logger
}
