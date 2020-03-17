package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/comfortablynumb/goginrestapi/internal/componentregistry"
	context2 "github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/errorhandler"
	hooks2 "github.com/comfortablynumb/goginrestapi/internal/hooks"
	"github.com/comfortablynumb/goginrestapi/internal/middleware"
	"github.com/comfortablynumb/goginrestapi/internal/module"
	"github.com/comfortablynumb/goginrestapi/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"gopkg.in/go-playground/validator.v9"

	_ "github.com/mattn/go-sqlite3"

	"github.com/comfortablynumb/goginrestapi/internal/config"
)

// Constants

const (
	DbDriverName = "sqlite3"
)

// Interfaces

type App interface {
	Run() error
}

// Structs

type app struct {
	config            *config.AppConfig
	componentRegistry *componentregistry.ComponentRegistry
	hooks             *hooks2.Hooks
	errorHandler      *errorhandler.ErrorHandler
	router            *gin.Engine
	logger            *zerolog.Logger
	translator        *ut.UniversalTranslator
	moduleManager     *module.ModuleManager
}

func (a *app) Run() error {
	a.setUp()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Port),
		Handler: a.router,
	}

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.errorHandler.HandleFatal(err, "There was an error while starting listening for incoming requests on the web server.")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.logger.Debug().Msg("[app] Shutdown Server. Waiting a maximum of 15 seconds to finish pending work...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		a.errorHandler.HandleFatal(err, "There was an error while shutting down the web server.")
	}

	a.logger.Debug().Msg("[app] Server exiting.")

	return nil
}

func (a *app) setUp() {
	a.logger = a.createLogger()
	a.translator = a.createTranslator()
	a.errorHandler = a.createErrorHandler()
	a.moduleManager = a.createModuleManager()
	a.componentRegistry = a.createComponentRegistry()
	a.router = a.createRouter()

	a.setUpValidator(a.componentRegistry.Validator)

	a.executeDbMigrations(a.componentRegistry.Db)
}

func (a *app) createLogger() *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	logger := zerolog.New(output).With().Timestamp().Logger()

	return a.hooks.SetupLogger(&logger)
}

func (a *app) createDb() *sql.DB {
	a.logger.Debug().Msgf("[app] Creating DB instance for driver: %s", DbDriverName)

	db, err := sql.Open(DbDriverName, a.config.DbUri)

	a.errorHandler.HandleFatalIfError(err, "Could NOT create a DB instance.")

	a.logger.Debug().Msg("[app] Executing ping on the database.")

	err = db.Ping()

	a.errorHandler.HandleFatalIfError(err, "There was an error while trying to ping the database.")

	return db
}

func (a *app) createModuleManager() *module.ModuleManager {
	moduleManager := module.NewModuleManager()

	moduleManager.AddModule(&module.UserTypeModule{})
	moduleManager.AddModule(&module.UserModule{})

	return moduleManager
}

func (a *app) setUpValidator(validator *validator.Validate) {
	for _, m := range a.moduleManager.GetModules() {
		a.logger.Debug().Msgf("[app] Setting up validator for module '%s'...", m.GetName())

		m.SetUpValidator(a.errorHandler, a.componentRegistry, a.componentRegistry.Validator)
	}
}

func (a *app) executeDbMigrations(db *sql.DB) {
	a.logger.Debug().Msg("[app] Creating database migrations driver.")

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	a.errorHandler.HandleFatalIfError(err, "Could NOT create database migrations driver.")

	a.logger.Debug().Msg("[app] Creating database migrations instance.")

	databaseMigrations, err := migrate.NewWithDatabaseInstance(
		a.config.DbMigrationsPath,
		DbDriverName,
		driver,
	)

	a.errorHandler.HandleFatalIfError(err, "Could NOT create database migrations instance.")

	a.logger.Debug().Msg("[app] Executing database migrations.")

	err = databaseMigrations.Up()

	a.errorHandler.HandleFatalIfError(err, "There was an error while trying to execute the database migrations.")

	a.logger.Debug().Msg("[app] Database migrations executed SUCCESSFULLY!")
}

func (a *app) createErrorHandler() *errorhandler.ErrorHandler {
	return errorhandler.NewErrorHandler(a.logger, a.hooks)
}

func (a *app) createValidator() *validator.Validate {
	return a.hooks.SetupValidator(validator.New())
}

func (a *app) createTranslator() *ut.UniversalTranslator {
	enLocale := en.New()
	esLocale := es.New()

	return ut.New(enLocale, esLocale)
}

func (a *app) createRequestContextFactory() *context2.RequestContextFactory {
	return context2.NewRequestContextFactory(a.translator)
}

func (a *app) createTimeService() service.TimeService {
	return service.NewTimeService()
}

func (a *app) createComponentRegistry() *componentregistry.ComponentRegistry {
	componentRegistry := componentregistry.NewComponentRegistry()

	// Logger

	componentRegistry.Logger = a.logger

	// Validator

	componentRegistry.Validator = a.createValidator()

	// Translator

	componentRegistry.Translator = a.translator

	// Request Context Factory

	componentRegistry.RequestContextFactory = a.createRequestContextFactory()

	// Time Service

	componentRegistry.TimeService = a.createTimeService()

	// Db

	componentRegistry.Db = a.createDb()

	// Register modules components

	for _, m := range a.moduleManager.GetModules() {
		a.logger.Debug().Msgf("[app] Registering module '%s' components...", m.GetName())

		m.SetUpComponents(a.errorHandler, componentRegistry)
	}

	return componentRegistry
}

func (a *app) createRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.ErrorHandler(a.componentRegistry.RequestContextFactory, gin.ErrorTypeAny, a.errorHandler))

	router = a.hooks.SetupRouter(router)

	// Setup modules routes

	for _, m := range a.moduleManager.GetModules() {
		a.logger.Debug().Msgf("[app] Registering module '%s' routes...", m.GetName())

		m.SetUpRouter(a.errorHandler, a.componentRegistry, router)
	}

	return router
}

// Static functions

func NewApp(appConfig *config.AppConfig) App {
	return &app{
		config: appConfig,
		hooks:  hooks2.NewHooks(),
	}
}

func NewAppFromEnv() (App, error) {
	appConfig := config.NewAppConfig()
	err := envconfig.Process("MYAPP", appConfig)

	if err != nil {
		return nil, err
	}

	return NewApp(appConfig), nil
}
