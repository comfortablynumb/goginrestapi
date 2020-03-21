package resource

import (
	"time"

	"github.com/comfortablynumb/goginrestapi/internal/model"
)

// Structs

// UserFindResource

type UserFindResource struct {
	CommonFindResource

	Username *string `form:"username" validate:"omitempty,min=1,max=50"`
}

// UserCreateResource

type UserCreateResource struct {
	Username     string `json:"username" binding:"required" validate:"required,min=1,max=50"`
	UserTypeName string `json:"user_type_name" validate:"required,user_type"`
	Disabled     bool   `json:"disabled"`
}

// UserUpdateResource

type UserUpdateResource struct {
	Username     string `uri:"username" json:"-" binding:"required" validate:"required,min=1,max=50"`
	UserTypeName string `json:"user_type_name" validate:"required,user_type"`
	Disabled     bool   `json:"disabled"`
}

// UserDeleteResource

type UserDeleteResource struct {
	Username string `uri:"username" json:"-" binding:"required" validate:"required,min=1,max=50"`
}

// UserResource

type UserResource struct {
	Username  string           `json:"username"`
	UserType  UserTypeResource `json:"user_type"`
	Disabled  bool             `json:"disabled"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

// UserResourceBuilder

type UserResourceBuilder struct {
	username  string
	userType  UserTypeResource
	disabled  bool
	createdAt time.Time
	updatedAt time.Time
}

func (b *UserResourceBuilder) WithUsername(username string) *UserResourceBuilder {
	b.username = username

	return b
}

func (b *UserResourceBuilder) WithUserType(userType model.UserType) *UserResourceBuilder {
	b.userType = *FromUserType(userType)

	return b
}

func (b *UserResourceBuilder) WithDisabled(disabled bool) *UserResourceBuilder {
	b.disabled = disabled

	return b
}

func (b *UserResourceBuilder) WithCreatedAt(createdAt time.Time) *UserResourceBuilder {
	b.createdAt = createdAt

	return b
}

func (b *UserResourceBuilder) WithUpdatedAt(updatedAt time.Time) *UserResourceBuilder {
	b.updatedAt = updatedAt

	return b
}

func (b *UserResourceBuilder) Build() *UserResource {
	return NewUserResource(b.username, b.userType, b.disabled, b.createdAt, b.updatedAt)
}

// Static functions

func NewUserResourceBuilder() *UserResourceBuilder {
	return &UserResourceBuilder{}
}

func NewUserResource(
	username string,
	userType UserTypeResource,
	disabled bool,
	createdAt time.Time,
	updatedAt time.Time,
) *UserResource {
	return &UserResource{
		Username:  username,
		UserType:  userType,
		Disabled:  disabled,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func FromUser(user model.User) *UserResource {
	return NewUserResourceBuilder().
		WithUsername(user.Username).
		WithUserType(user.UserType).
		WithDisabled(user.Disabled).
		WithCreatedAt(user.CreatedAt).
		WithUpdatedAt(user.UpdatedAt).
		Build()
}
