package module

import (
	"github.com/comfortablynumb/goginrestapi/internal/componentregistry"
	"github.com/comfortablynumb/goginrestapi/internal/controller"
	"github.com/comfortablynumb/goginrestapi/internal/errorhandler"
	repository2 "github.com/comfortablynumb/goginrestapi/internal/repository"
	"github.com/comfortablynumb/goginrestapi/internal/service"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// Constants

const (
	UserModuleName              = "user"
	UserRepositoryComponentName = "UserRepository"
	UserServiceComponentName    = "UserService"
	UserControllerComponentName = "UserController"
)

// Structs

type UserModule struct {
}

func (m *UserModule) GetName() string {
	return UserModuleName
}

func (m *UserModule) SetUpComponents(errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry) {
	userTypeService := componentRegistry.GetOrPanic(UserTypeServiceComponentName).(service.UserTypeService)

	repo := repository2.NewUserRepository(componentRegistry.Db, componentRegistry.Logger)
	serv := service.NewUserService(
		componentRegistry.Logger,
		componentRegistry.Validator,
		componentRegistry.TimeService,
		repo,
		userTypeService,
	)
	cont := controller.NewUserController(serv, componentRegistry.RequestContextFactory)

	componentRegistry.Set(UserRepositoryComponentName, repo).
		Set(UserServiceComponentName, serv).
		Set(UserControllerComponentName, cont)
}

func (m *UserModule) SetUpRouter(errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry, router *gin.Engine) {
	userController := componentRegistry.GetOrPanic(UserControllerComponentName).(*controller.UserController)

	users := router.Group("/user")

	users.GET("", userController.Find)
	users.POST("", userController.Create)
	users.PUT("/:username", userController.Update)
	users.DELETE("/:username", userController.Delete)
}

func (m *UserModule) SetUpValidator(errorHandler *errorhandler.ErrorHandler, componentRegistry *componentregistry.ComponentRegistry, validator *validator.Validate) {

}
