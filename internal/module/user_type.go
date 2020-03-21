package module

import (
	"github.com/comfortablynumb/goginrestapi/internal/componentregistry"
	"github.com/comfortablynumb/goginrestapi/internal/config"
	"github.com/comfortablynumb/goginrestapi/internal/controller"
	"github.com/comfortablynumb/goginrestapi/internal/errorhandler"
	"github.com/comfortablynumb/goginrestapi/internal/repository"
	"github.com/comfortablynumb/goginrestapi/internal/service"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// Constants

const (
	UserTypeModuleName              = "user_type"
	UserTypeRepositoryComponentName = "UserTypeRepository"
	UserTypeServiceComponentName    = "UserTypeService"
	UserTypeControllerComponentName = "UserTypeController"
)

// Structs

type UserTypeModule struct {
}

func (m *UserTypeModule) GetName() string {
	return UserTypeModuleName
}

func (m *UserTypeModule) SetUpComponents(
	appConfig config.AppConfig,
	errorHandler *errorhandler.ErrorHandler,
	componentRegistry *componentregistry.ComponentRegistry,
) {
	repo := repository.NewUserTypeRepository(appConfig, componentRegistry.Db, componentRegistry.Logger)
	serv := service.NewUserTypeService(
		appConfig,
		componentRegistry.Logger,
		componentRegistry.Validator,
		componentRegistry.TimeService,
		repo,
	)
	cont := controller.NewUserTypeController(serv, componentRegistry.RequestContextFactory)

	componentRegistry.Set(UserTypeRepositoryComponentName, repo).
		Set(UserTypeServiceComponentName, serv).
		Set(UserTypeControllerComponentName, cont)
}

func (m *UserTypeModule) SetUpRouter(errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry, router *gin.Engine) {
	userTypeController := componentRegistry.GetOrPanic(UserTypeControllerComponentName).(*controller.UserTypeController)

	userTypes := router.Group("/user_type")

	userTypes.GET("", userTypeController.Find)
	userTypes.POST("", userTypeController.Create)
	userTypes.PUT("/:name", userTypeController.Update)
	userTypes.DELETE("/:name", userTypeController.Delete)
}

func (m *UserTypeModule) SetUpValidator(errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry, validator *validator.Validate) {
	userTypeService := componentRegistry.GetOrPanic(UserTypeServiceComponentName).(service.UserTypeService)

	errorHandler.HandleFatalIfError(
		validator.RegisterValidationCtx("user_type", userTypeService.ValidateUserTypeByName),
		"Could NOT register user type validation.",
	)
}
