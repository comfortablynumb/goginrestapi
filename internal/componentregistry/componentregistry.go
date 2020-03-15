package componentregistry

import (
	"database/sql"

	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/controller"
	repository2 "github.com/comfortablynumb/goginrestapi/internal/repository"
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

	UserTypeController *controller.UserTypeController
	UserTypeService    service.UserTypeService
	UserTypeRepository repository2.UserTypeRepository

	UserController *controller.UserController
	UserService    service.UserService
	UserRepository repository2.UserRepository
}

// Static functions

func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{}
}
