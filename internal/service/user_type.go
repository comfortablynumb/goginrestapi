package service

import (
	context2 "context"

	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/comfortablynumb/goginrestapi/internal/config"
	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/model"
	"github.com/comfortablynumb/goginrestapi/internal/repository"
	"github.com/comfortablynumb/goginrestapi/internal/repository/utils"
	"github.com/comfortablynumb/goginrestapi/internal/resource"
	"github.com/rs/zerolog"
	validator2 "gopkg.in/go-playground/validator.v9"
)

// Constants

const (
	UserTypeServiceSourceName = "UserTypeService"
)

// Interfaces

type UserTypeService interface {
	Count(ctx *context.RequestContext, userTypeFindResource *resource.UserTypeFindResource) (int64, *apperror.AppError)
	Find(ctx *context.RequestContext, userTypeFindResource *resource.UserTypeFindResource) (*resource.UserTypeResourceList, *apperror.AppError)
	FindOneByName(ctx *context.RequestContext, name string) (*resource.UserTypeResource, *apperror.AppError)
	Create(ctx *context.RequestContext, userCreateResource *resource.UserTypeCreateResource) (*resource.UserTypeResource, *apperror.AppError)
	Update(ctx *context.RequestContext, userUpdateResource *resource.UserTypeUpdateResource) (*resource.UserTypeResource, *apperror.AppError)
	Delete(ctx *context.RequestContext, userDeleteResource *resource.UserTypeDeleteResource) (*resource.UserTypeResource, *apperror.AppError)
	ValidateUserTypeByName(ctx context2.Context, fl validator2.FieldLevel) bool
	ValidateUserTypeUnique(ctx context2.Context, sl validator2.StructLevel)
}

// Structs

type userTypeService struct {
	appConfig          config.AppConfig
	logger             *zerolog.Logger
	validator          *validator2.Validate
	timeService        TimeService
	userTypeRepository repository.UserTypeRepository
}

func (s *userTypeService) Count(ctx *context.RequestContext, userTypeFindResource *resource.UserTypeFindResource) (int64, *apperror.AppError) {
	filters := utils.NewUserTypeFindFiltersBuilder().WithName(userTypeFindResource.Name).Build()
	options := utils.NewUserTypeFindOptionsBuilder(s.appConfig.DefaultLimit).
		WithSortBy(userTypeFindResource.SortBy, userTypeFindResource.SortDir).
		WithLimit(userTypeFindResource.Offset, userTypeFindResource.Limit).
		Build()

	return s.userTypeRepository.Count(ctx, filters, options)
}

func (s *userTypeService) Find(ctx *context.RequestContext, userTypeFindResource *resource.UserTypeFindResource) (*resource.UserTypeResourceList, *apperror.AppError) {
	if err := s.validator.StructCtx(ctx, userTypeFindResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	count, err := s.Count(ctx, userTypeFindResource)

	if err != nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, UserTypeServiceSourceName)
	}

	result := make([]*resource.UserTypeResource, 0)

	if count < 1 {
		return resource.NewUserTypeResourceList(result, count), nil
	}

	filters := utils.NewUserTypeFindFiltersBuilder().WithName(userTypeFindResource.Name).Build()
	options := utils.NewUserTypeFindOptionsBuilder(s.appConfig.DefaultLimit).
		WithSortBy(userTypeFindResource.SortBy, userTypeFindResource.SortDir).
		WithLimit(userTypeFindResource.Offset, userTypeFindResource.Limit).
		Build()

	rows, err := s.userTypeRepository.Find(ctx, filters, options)

	if err != nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, UserTypeServiceSourceName)
	}

	for _, row := range rows {
		result = append(result, resource.FromUserType(*row))
	}

	return resource.NewUserTypeResourceList(result, count), nil
}

func (s *userTypeService) FindOneByName(ctx *context.RequestContext, name string) (*resource.UserTypeResource, *apperror.AppError) {
	if err := s.validator.VarCtx(ctx, name, "required"); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	userType, err := s.userTypeRepository.FindOneByName(ctx, name)

	if err != nil {
		return nil, err
	}

	if userType == nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, UserTypeServiceSourceName)
	}

	return resource.FromUserType(*userType), nil
}

func (s *userTypeService) Create(ctx *context.RequestContext, userCreateResource *resource.UserTypeCreateResource) (*resource.UserTypeResource, *apperror.AppError) {
	if err := s.validator.StructCtx(ctx, userCreateResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	userType := model.NewUserTypeBuilder().
		WithName(userCreateResource.Name).
		WithDisabled(userCreateResource.Disabled).
		WithCreatedAt(s.timeService.GetCurrentUtcTime()).
		WithUpdatedAt(s.timeService.GetCurrentUtcTime()).
		Build()

	err := s.userTypeRepository.Create(ctx, userType)

	if err != nil {
		return nil, err
	}

	return resource.FromUserType(*userType), nil
}

func (s *userTypeService) Update(ctx *context.RequestContext, userUpdateResource *resource.UserTypeUpdateResource) (*resource.UserTypeResource, *apperror.AppError) {
	userType, err := s.userTypeRepository.FindOneByName(ctx, userUpdateResource.OriginalName)

	if err != nil {
		return nil, err
	}

	if userType == nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, UserTypeServiceSourceName)
	}

	userUpdateResource.ID = userType.ID

	if err := s.validator.StructCtx(ctx, userUpdateResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	userType.Name = userUpdateResource.Name
	userType.Disabled = userUpdateResource.Disabled
	userType.UpdatedAt = s.timeService.GetCurrentUtcTime()

	err = s.userTypeRepository.Update(ctx, userType)

	if err != nil {
		return nil, err
	}

	return resource.FromUserType(*userType), nil
}

func (s *userTypeService) Delete(ctx *context.RequestContext, userTypeDeleteResource *resource.UserTypeDeleteResource) (*resource.UserTypeResource, *apperror.AppError) {
	if err := s.validator.StructCtx(ctx, userTypeDeleteResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	userType, err := s.userTypeRepository.FindOneByName(ctx, userTypeDeleteResource.Name)

	if err != nil {
		return nil, err
	}

	if userType == nil {
		return nil, nil
	}

	err = s.userTypeRepository.Delete(ctx, userType)

	if err != nil {
		return nil, err
	}

	return resource.FromUserType(*userType), nil
}

func (s *userTypeService) ValidateUserTypeByName(ctx context2.Context, fl validator2.FieldLevel) bool {
	requestCtx := ctx.(*context.RequestContext)
	userTypeName := fl.Field().String()
	userType, err := s.userTypeRepository.FindOneByName(requestCtx, userTypeName)

	if err != nil {
		s.logger.Err(err)

		return false
	}

	if userType == nil {
		return false
	}

	requestCtx.Set("user_type", userType)

	return true
}

func (s *userTypeService) ValidateUserTypeUnique(ctx context2.Context, sl validator2.StructLevel) {
	requestCtx := ctx.(*context.RequestContext)
	userType := sl.Current().Interface().(resource.UserTypeUniqueValidator)

	if len(userType.GetName()) > 0 {
		currentUserType, err := s.userTypeRepository.FindOneByName(requestCtx, userType.GetName())

		if err != nil {
			s.logger.Err(err)

			sl.ReportError(userType.GetName(), "Name", "Name", "unique", "")

			return
		}

		if currentUserType != nil && currentUserType.ID != userType.GetID() {
			sl.ReportError(userType.GetName(), "Name", "Name", "unique", "")

			return
		}
	}
}

// Static functions

func NewUserTypeService(
	appConfig config.AppConfig,
	logger *zerolog.Logger,
	validator *validator2.Validate,
	timeService TimeService,
	userTypeRepository repository.UserTypeRepository,
) UserTypeService {
	return &userTypeService{
		appConfig:          appConfig,
		logger:             logger,
		validator:          validator,
		timeService:        timeService,
		userTypeRepository: userTypeRepository,
	}
}
