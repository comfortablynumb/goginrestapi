package resource

import (
	"time"

	"github.com/comfortablynumb/goginrestapi/internal/model"
)

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
	Name      string    `json:"name"`
	Disabled  bool      `json:"disabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserTypeResourceBuilder struct {
	name      string
	disabled  bool
	createdAt time.Time
	updatedAt time.Time
}

func (b *UserTypeResourceBuilder) WithName(name string) *UserTypeResourceBuilder {
	b.name = name

	return b
}

func (b *UserTypeResourceBuilder) WithDisabled(disabled bool) *UserTypeResourceBuilder {
	b.disabled = disabled

	return b
}

func (b *UserTypeResourceBuilder) WithCreatedAt(createdAt time.Time) *UserTypeResourceBuilder {
	b.createdAt = createdAt

	return b
}

func (b *UserTypeResourceBuilder) WithUpdatedAt(updatedAt time.Time) *UserTypeResourceBuilder {
	b.updatedAt = updatedAt

	return b
}

func (b *UserTypeResourceBuilder) Build() *UserTypeResource {
	return NewUserTypeResource(b.name, b.disabled, b.createdAt, b.updatedAt)
}

// Static functions

func NewUserTypeResourceBuilder() *UserTypeResourceBuilder {
	return &UserTypeResourceBuilder{}
}

func NewUserTypeResource(name string, disabled bool, createdAt time.Time, updatedAt time.Time) *UserTypeResource {
	return &UserTypeResource{
		Name:      name,
		Disabled:  disabled,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func FromUserType(userType model.UserType) *UserTypeResource {
	return NewUserTypeResourceBuilder().
		WithName(userType.Name).
		WithDisabled(userType.Disabled).
		WithCreatedAt(userType.CreatedAt).
		WithUpdatedAt(userType.UpdatedAt).
		Build()
}
