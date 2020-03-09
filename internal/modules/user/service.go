package user

import (
	"github.com/comfortablynumb/goginrestapi/internal/apperror"
	"github.com/rs/zerolog"
	validator2 "gopkg.in/go-playground/validator.v9"
)

// Constants

const (
	ServiceSourceName = "UserService"
)

// Interfaces

type UserService interface {
	Find(userFindResource *UserFindResource) ([]*UserResource, *apperror.AppError)
	Create(userCreateResource *UserCreateResource) (*UserResource, *apperror.AppError)
	Update(userUpdateResource *UserUpdateResource) (*UserResource, *apperror.AppError)
	Delete(userDeleteResource *UserDeleteResource) (*UserResource, *apperror.AppError)
}

// Structs

type userService struct {
	validator      *validator2.Validate
	userRepository UserRepository
	logger         *zerolog.Logger
}

func (s *userService) Find(userFindResource *UserFindResource) ([]*UserResource, *apperror.AppError) {
	if err := s.validator.Struct(userFindResource); err != nil {
		return nil, apperror.NewValidationAppError(err, ServiceSourceName)
	}

	filters := NewUserFindFilters().WithUsernamePtr(userFindResource.Username)
	options := NewUserFindOptions().
		WithSortByPtr(userFindResource.SortBy, userFindResource.SortDir).
		WithLimitPtr(userFindResource.Offset, userFindResource.Limit)

	rows, err := s.userRepository.Find(filters, options)

	if err != nil {
		return nil, apperror.NewModelNotFoundAppError(err, ServiceSourceName)
	}

	result := make([]*UserResource, 0)

	for _, row := range rows {
		result = append(result, NewUserResource(row.Username, row.Disabled))
	}

	return result, nil
}

func (s *userService) Create(userCreateResource *UserCreateResource) (*UserResource, *apperror.AppError) {
	if err := s.validator.Struct(userCreateResource); err != nil {
		return nil, apperror.NewValidationAppError(err, ServiceSourceName)
	}

	user := NewUser(0, userCreateResource.Username, userCreateResource.Disabled)

	err := s.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return NewUserResource(user.Username, user.Disabled), nil
}

func (s *userService) Update(userUpdateResource *UserUpdateResource) (*UserResource, *apperror.AppError) {
	if err := s.validator.Struct(userUpdateResource); err != nil {
		return nil, apperror.NewValidationAppError(err, ServiceSourceName)
	}

	user, err := s.userRepository.FindOneByUsername(userUpdateResource.Username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, apperror.NewModelNotFoundAppError(err, ServiceSourceName)
	}

	user.Username = userUpdateResource.Username
	user.Disabled = userUpdateResource.Disabled

	err = s.userRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return NewUserResource(user.Username, user.Disabled), nil
}

func (s *userService) Delete(userDeleteResource *UserDeleteResource) (*UserResource, *apperror.AppError) {
	if err := s.validator.Struct(userDeleteResource); err != nil {
		return nil, apperror.NewValidationAppError(err, ServiceSourceName)
	}

	user, err := s.userRepository.FindOneByUsername(userDeleteResource.Username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	err = s.userRepository.Delete(user)

	if err != nil {
		return nil, err
	}

	return NewUserResource(user.Username, user.Disabled), nil
}

// Static functions

func NewUserService(validator *validator2.Validate, userRepository UserRepository, logger *zerolog.Logger) UserService {
	return &userService{
		validator:      validator,
		userRepository: userRepository,
		logger:         logger,
	}
}
