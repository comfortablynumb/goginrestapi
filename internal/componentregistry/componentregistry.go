package componentregistry

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/service"
	ut "github.com/go-playground/universal-translator"
	"github.com/rs/zerolog"
	"gopkg.in/go-playground/validator.v9"
)

// Structs

type ComponentRegistry struct {
	Db                    *sql.DB
	Validator             *validator.Validate
	Logger                *zerolog.Logger
	Translator            *ut.UniversalTranslator
	RequestContextFactory *context.RequestContextFactory

	TimeService service.TimeService

	Components map[string]interface{}
}

func (c *ComponentRegistry) Set(name string, component interface{}) *ComponentRegistry {
	c.Components[name] = component

	return c
}

func (c *ComponentRegistry) Get(name string) (interface{}, error) {
	component, found := c.Components[name]

	if !found {
		return nil, errors.New(fmt.Sprintf("Component '%s' is not registered in the component registry.", name))
	}

	return component, nil
}

func (c *ComponentRegistry) GetOrPanic(name string) interface{} {
	component, err := c.Get(name)

	if err != nil {
		panic(err)
	}

	return component
}

// Static functions

func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{
		Components: make(map[string]interface{}),
	}
}
