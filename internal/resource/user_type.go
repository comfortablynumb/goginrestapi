package resource

// Structs

type UserTypeFindResource struct {
	Name    *string `form:"name" validate:"omitempty,min=1,max=50"`
	SortBy  *string `form:"sort_by"`
	SortDir *string `form:"sort_dir"`
	Offset  *int    `form:"offset"`
	Limit   *int    `form:"limit"`
}

type UserTypeCreateResource struct {
	Name     string `json:"name" binding:"required" validate:"required,min=1,max=50"`
	Disabled bool   `json:"disabled"`
}

type UserTypeUpdateResource struct {
	Name     string `uri:"name" binding:"required" validate:"required,min=1,max=50"`
	Disabled bool   `json:"disabled"`
}

type UserTypeDeleteResource struct {
	Name string `uri:"name" binding:"required" validate:"required,min=1,max=50"`
}

type UserTypeResource struct {
	Name     string `json:"name"`
	Disabled bool   `json:"disabled"`
}

// Static functions

func NewUserTypeResource(name string, disabled bool) *UserTypeResource {
	return &UserTypeResource{
		Name:     name,
		Disabled: disabled,
	}
}
