package resource

// Structs

type UserFindResource struct {
	Username *string `form:"username" validate:"omitempty,min=1,max=50"`
	SortBy   *string `form:"sort_by"`
	SortDir  *string `form:"sort_dir"`
	Offset   *int    `form:"offset"`
	Limit    *int    `form:"limit"`
}

type UserCreateResource struct {
	Username string `json:"username" binding:"required" validate:"required,min=1,max=50"`
	Disabled bool   `json:"disabled"`
}

type UserUpdateResource struct {
	Username string `uri:"username" binding:"required" validate:"required,min=1,max=50"`
	Disabled bool   `json:"disabled"`
}

type UserDeleteResource struct {
	Username string `uri:"username" binding:"required" validate:"required,min=1,max=50"`
}

type UserResource struct {
	Username string `json:"username"`
	Disabled bool   `json:"disabled"`
}

// Static functions

func NewUserResource(username string, disabled bool) *UserResource {
	return &UserResource{
		Username: username,
		Disabled: disabled,
	}
}
