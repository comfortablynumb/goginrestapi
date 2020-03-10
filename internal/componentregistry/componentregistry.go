package componentregistry

import (
	"database/sql"

	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/modules/user"
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

	UserController *user.UserController
	UserService    user.UserService
	UserRepository user.UserRepository
}

// Static functions

func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{}
}
