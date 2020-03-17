package service

import (
	context2 "context"

	"github.com/comfortablynumb/goginrestapi/internal/apperror"
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
	Find(ctx *context.RequestContext, userTypeFindResource *resource.UserTypeFindResource) ([]*resource.UserTypeResource, *apperror.AppError)
	Create(ctx *context.RequestContext, userCreateResource *resource.UserTypeCreateResource) (*resource.UserTypeResource, *apperror.AppError)
	Update(ctx *context.RequestContext, userUpdateResource *resource.UserTypeUpdateResource) (*resource.UserTypeResource, *apperror.AppError)
	Delete(ctx *context.RequestContext, userDeleteResource *resource.UserTypeDeleteResource) (*resource.UserTypeResource, *apperror.AppError)
	ValidateUserTypeByName(ctx context2.Context, fl validator2.FieldLevel) bool
}

// Structs

type userTypeService struct {
	logger             *zerolog.Logger
	validator          *validator2.Validate
	timeService        TimeService
	userTypeRepository repository.UserTypeRepository
}

func (s *userTypeService) Find(ctx *context.RequestContext, userTypeFindResource *resource.UserTypeFindResource) ([]*resource.UserTypeResource, *apperror.AppError) {
	if err := s.validator.StructCtx(ctx, userTypeFindResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	filters := utils.NewUserTypeFindFilters().WithNamePtr(userTypeFindResource.Name)
	options := utils.NewUserTypeFindOptions().
		WithSortByPtr(userTypeFindResource.SortBy, userTypeFindResource.SortDir).
		WithLimitPtr(userTypeFindResource.Offset, userTypeFindResource.Limit)

	rows, err := s.userTypeRepository.Find(ctx, filters, options)

	if err != nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, UserTypeServiceSourceName)
	}

	result := make([]*resource.UserTypeResource, 0)

	for _, row := range rows {
		result = append(result, resource.FromUserType(*row))
	}

	return result, nil
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
	if err := s.validator.StructCtx(ctx, userUpdateResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	userType, err := s.userTypeRepository.FindOneByName(ctx, userUpdateResource.Name)

	if err != nil {
		return nil, err
	}

	if userType == nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, UserTypeServiceSourceName)
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

func (s *userTypeService) Delete(ctx *context.RequestContext, userDeleteResource *resource.UserTypeDeleteResource) (*resource.UserTypeResource, *apperror.AppError) {
	if err := s.validator.StructCtx(ctx, userDeleteResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	userType, err := s.userTypeRepository.FindOneByName(ctx, userDeleteResource.Name)

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

// Static functions

func NewUserTypeService(
	logger *zerolog.Logger,
	validator *validator2.Validate,
	timeService TimeService,
	userTypeRepository repository.UserTypeRepository,
) UserTypeService {
	return &userTypeService{
		logger:             logger,
		validator:          validator,
		timeService:        timeService,
		userTypeRepository: userTypeRepository,
	}
}
