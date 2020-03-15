package user

import (
	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/services"
	"github.com/rs/zerolog"
	validator2 "gopkg.in/go-playground/validator.v9"
)

// Constants

const (
	ServiceSourceName = "UserService"
)

// Interfaces

type UserService interface {
	Find(ctx *context.RequestContext, userFindResource *UserFindResource) ([]*UserResource, *apperror.AppError)
	Create(ctx *context.RequestContext, userCreateResource *UserCreateResource) (*UserResource, *apperror.AppError)
	Update(ctx *context.RequestContext, userUpdateResource *UserUpdateResource) (*UserResource, *apperror.AppError)
	Delete(ctx *context.RequestContext, userDeleteResource *UserDeleteResource) (*UserResource, *apperror.AppError)
}

// Structs

type userService struct {
	logger         *zerolog.Logger
	validator      *validator2.Validate
	timeService    services.TimeService
	userRepository UserRepository
}

func (s *userService) Find(ctx *context.RequestContext, userFindResource *UserFindResource) ([]*UserResource, *apperror.AppError) {
	if err := s.validator.Struct(userFindResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, ServiceSourceName)
	}

	filters := NewUserFindFilters().WithUsernamePtr(userFindResource.Username)
	options := NewUserFindOptions().
		WithSortByPtr(userFindResource.SortBy, userFindResource.SortDir).
		WithLimitPtr(userFindResource.Offset, userFindResource.Limit)

	rows, err := s.userRepository.Find(ctx, filters, options)

	if err != nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, ServiceSourceName)
	}

	result := make([]*UserResource, 0)

	for _, row := range rows {
		result = append(result, NewUserResource(row.Username, row.Disabled))
	}

	return result, nil
}

func (s *userService) Create(ctx *context.RequestContext, userCreateResource *UserCreateResource) (*UserResource, *apperror.AppError) {
	if err := s.validator.Struct(userCreateResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, ServiceSourceName)
	}

	user := NewUserBuilder().
		WithUsername(userCreateResource.Username).
		WithDisabled(userCreateResource.Disabled).
		WithCreatedAt(s.timeService.GetCurrentUtcTime()).
		WithUpdatedAt(s.timeService.GetCurrentUtcTime()).
		Build()

	err := s.userRepository.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return NewUserResource(user.Username, user.Disabled), nil
}

func (s *userService) Update(ctx *context.RequestContext, userUpdateResource *UserUpdateResource) (*UserResource, *apperror.AppError) {
	if err := s.validator.Struct(userUpdateResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, ServiceSourceName)
	}

	user, err := s.userRepository.FindOneByUsername(ctx, userUpdateResource.Username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, ServiceSourceName)
	}

	user.Username = userUpdateResource.Username
	user.Disabled = userUpdateResource.Disabled
	user.UpdatedAt = s.timeService.GetCurrentUtcTime()

	err = s.userRepository.Update(ctx, user)

	if err != nil {
		return nil, err
	}

	return NewUserResource(user.Username, user.Disabled), nil
}

func (s *userService) Delete(ctx *context.RequestContext, userDeleteResource *UserDeleteResource) (*UserResource, *apperror.AppError) {
	if err := s.validator.Struct(userDeleteResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, ServiceSourceName)
	}

	user, err := s.userRepository.FindOneByUsername(ctx, userDeleteResource.Username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	err = s.userRepository.Delete(ctx, user)

	if err != nil {
		return nil, err
	}

	return NewUserResource(user.Username, user.Disabled), nil
}

// Static functions

func NewUserService(
	logger *zerolog.Logger,
	validator *validator2.Validate,
	timeService services.TimeService,
	userRepository UserRepository,
) UserService {
	return &userService{
		logger:         logger,
		validator:      validator,
		timeService:    timeService,
		userRepository: userRepository,
	}
}
