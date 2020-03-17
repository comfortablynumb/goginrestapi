package resource

import (
	"time"

	"github.com/comfortablynumb/goginrestapi/internal/model"
)

// Structs

type UserFindResource struct {
	Username *string `form:"username" validate:"omitempty,min=1,max=50"`
	SortBy   *string `form:"sort_by"`
	SortDir  *string `form:"sort_dir"`
	Offset   *int    `form:"offset"`
	Limit    *int    `form:"limit"`
}

type UserCreateResource struct {
	Username     string `json:"username" binding:"required" validate:"required,min=1,max=50"`
	UserTypeName string `json:"user_type_name" validate:"required,user_type"`
	Disabled     bool   `json:"disabled"`
}

type UserUpdateResource struct {
	Username     string `uri:"username" binding:"required" validate:"required,min=1,max=50"`
	UserTypeName string `json:"user_type_name" validate:"required,user_type"`
	Disabled     bool   `json:"disabled"`
}

type UserDeleteResource struct {
	Username string `uri:"username" binding:"required" validate:"required,min=1,max=50"`
}

type UserResource struct {
	Username  string           `json:"username"`
	UserType  UserTypeResource `json:"user_type"`
	Disabled  bool             `json:"disabled"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

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
