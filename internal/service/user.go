package service

import (
	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/comfortablynumb/goginrestapi/internal/context"
	"github.com/comfortablynumb/goginrestapi/internal/model"
	repository2 "github.com/comfortablynumb/goginrestapi/internal/repository"
	"github.com/comfortablynumb/goginrestapi/internal/repository/utils"
	"github.com/comfortablynumb/goginrestapi/internal/resource"
	"github.com/rs/zerolog"
	validator2 "gopkg.in/go-playground/validator.v9"
)

// Constants

const (
	UserServiceSourceName = "UserService"
)

// Interfaces

type UserService interface {
	Find(ctx *context.RequestContext, userFindResource *resource.UserFindResource) ([]*resource.UserResource, *apperror.AppError)
	Create(ctx *context.RequestContext, userCreateResource *resource.UserCreateResource) (*resource.UserResource, *apperror.AppError)
	Update(ctx *context.RequestContext, userUpdateResource *resource.UserUpdateResource) (*resource.UserResource, *apperror.AppError)
	Delete(ctx *context.RequestContext, userDeleteResource *resource.UserDeleteResource) (*resource.UserResource, *apperror.AppError)
}

// Structs

type userService struct {
	logger          *zerolog.Logger
	validator       *validator2.Validate
	timeService     TimeService
	userRepository  repository2.UserRepository
	userTypeService UserTypeService
}

func (s *userService) Find(ctx *context.RequestContext, userFindResource *resource.UserFindResource) ([]*resource.UserResource, *apperror.AppError) {
	if err := s.validator.StructCtx(ctx, userFindResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserServiceSourceName)
	}

	filters := utils.NewUserFindFilters().WithUsernamePtr(userFindResource.Username)
	options := utils.NewUserFindOptions().
		WithSortByPtr(userFindResource.SortBy, userFindResource.SortDir).
		WithLimitPtr(userFindResource.Offset, userFindResource.Limit)

	rows, err := s.userRepository.Find(ctx, filters, options)

	if err != nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, UserServiceSourceName)
	}

	result := make([]*resource.UserResource, 0)

	for _, row := range rows {
		result = append(result, resource.FromUser(*row))
	}

	return result, nil
}

func (s *userService) Create(ctx *context.RequestContext, userCreateResource *resource.UserCreateResource) (*resource.UserResource, *apperror.AppError) {
	if err := s.validator.StructCtx(ctx, userCreateResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserServiceSourceName)
	}

	userType := ctx.Get("user_type").(*model.UserType)

	user := model.NewUserBuilder().
		WithUsername(userCreateResource.Username).
		WithUserType(*userType).
		WithDisabled(userCreateResource.Disabled).
		WithCreatedAt(s.timeService.GetCurrentUtcTime()).
		WithUpdatedAt(s.timeService.GetCurrentUtcTime()).
		Build()

	err := s.userRepository.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return resource.FromUser(*user), nil
}

func (s *userService) Update(ctx *context.RequestContext, userUpdateResource *resource.UserUpdateResource) (*resource.UserResource, *apperror.AppError) {
	if err := s.validator.StructCtx(ctx, userUpdateResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserServiceSourceName)
	}

	user, err := s.userRepository.FindOneByUsername(ctx, userUpdateResource.Username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, apperror.NewModelNotFoundAppError(ctx, err, UserServiceSourceName)
	}

	userType := ctx.Get("user_type").(*model.UserType)

	user.Username = userUpdateResource.Username
	user.UserType = *userType
	user.Disabled = userUpdateResource.Disabled
	user.UpdatedAt = s.timeService.GetCurrentUtcTime()

	err = s.userRepository.Update(ctx, user)

	if err != nil {
		return nil, err
	}

	return resource.FromUser(*user), nil
}

func (s *userService) Delete(ctx *context.RequestContext, userDeleteResource *resource.UserDeleteResource) (*resource.UserResource, *apperror.AppError) {
	if err := s.validator.StructCtx(ctx, userDeleteResource); err != nil {
		return nil, apperror.NewValidationAppError(ctx, err, UserServiceSourceName)
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

	return resource.FromUser(*user), nil
}

// Static functions

func NewUserService(
	logger *zerolog.Logger,
	validator *validator2.Validate,
	timeService TimeService,
	userRepository repository2.UserRepository,
	userTypeService UserTypeService,
) UserService {
	return &userService{
		logger:          logger,
		validator:       validator,
		timeService:     timeService,
		userRepository:  userRepository,
		userTypeService: userTypeService,
	}
}
