package resource

import "github.com/comfortablynumb/goginrestapi/internal/model"

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
	Username string           `json:"username"`
	UserType UserTypeResource `json:"user_type"`
	Disabled bool             `json:"disabled"`
}

// Static functions

func NewUserResource(username string, userType *UserTypeResource, disabled bool) *UserResource {
	return &UserResource{
		Username: username,
		UserType: *userType,
		Disabled: disabled,
	}
}

func FromUser(user *model.User) *UserResource {
	return NewUserResource(user.Username, FromUserType(&user.UserType), user.Disabled)
}
