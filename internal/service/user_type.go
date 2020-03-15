package service

import (
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
}

// Structs

type userTypeService struct {
	logger             *zerolog.Logger
	validator          *validator2.Validate
	timeService        TimeService
	userTypeRepository repository.UserTypeRepository
}

func (s *userTypeService) Find(ctx *context.RequestContext, userTypeFindResource *resource.UserTypeFindResource) ([]*resource.UserTypeResource, *apperror.AppError) {
	if err := s.validator.Struct(userTypeFindResource); err != nil {
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
		result = append(result, resource.NewUserTypeResource(row.Name, row.Disabled))
	}

	return result, nil
}

func (s *userTypeService) Create(ctx *context.RequestContext, userCreateResource *resource.UserTypeCreateResource) (*resource.UserTypeResource, *apperror.AppError) {
	if err := s.validator.Struct(userCreateResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	user := model.NewUserTypeBuilder().
		WithName(userCreateResource.Name).
		WithDisabled(userCreateResource.Disabled).
		WithCreatedAt(s.timeService.GetCurrentUtcTime()).
		WithUpdatedAt(s.timeService.GetCurrentUtcTime()).
		Build()

	err := s.userTypeRepository.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return resource.NewUserTypeResource(user.Name, user.Disabled), nil
}

func (s *userTypeService) Update(ctx *context.RequestContext, userUpdateResource *resource.UserTypeUpdateResource) (*resource.UserTypeResource, *apperror.AppError) {
	if err := s.validator.Struct(userUpdateResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	user, err := s.userTypeRepository.FindOneByName(ctx, userUpdateResource.Name)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, UserTypeServiceSourceName)
	}

	user.Name = userUpdateResource.Name
	user.Disabled = userUpdateResource.Disabled
	user.UpdatedAt = s.timeService.GetCurrentUtcTime()

	err = s.userTypeRepository.Update(ctx, user)

	if err != nil {
		return nil, err
	}

	return resource.NewUserTypeResource(user.Name, user.Disabled), nil
}

func (s *userTypeService) Delete(ctx *context.RequestContext, userDeleteResource *resource.UserTypeDeleteResource) (*resource.UserTypeResource, *apperror.AppError) {
	if err := s.validator.Struct(userDeleteResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserTypeServiceSourceName)
	}

	user, err := s.userTypeRepository.FindOneByName(ctx, userDeleteResource.Name)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	err = s.userTypeRepository.Delete(ctx, user)

	if err != nil {
		return nil, err
	}

	return resource.NewUserTypeResource(user.Name, user.Disabled), nil
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
