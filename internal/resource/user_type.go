package resource

import (
	"time"

	"github.com/comfortablynumb/goginrestapi/internal/model"
)

// Interfaces

type UserTypeUniqueValidator interface {
	GetID() int64
	GetName() string
}

// Structs

// UserTypeFindResource

type UserTypeFindResource struct {
	CommonFindResource

	Name *string `form:"name" validate:"omitempty,min=1,max=50"`
}

// UserTypeCreateResource

type UserTypeCreateResource struct {
	ID       int64  `json:"-"`
	Name     string `json:"name" binding:"required" validate:"required,min=1,max=50"`
	Disabled bool   `json:"disabled"`
}

func (u UserTypeCreateResource) GetID() int64 {
	return u.ID
}

func (u UserTypeCreateResource) GetName() string {
	return u.Name
}

// UserTypeUpdateResource

type UserTypeUpdateResource struct {
	ID           int64  `json:"-"`
	OriginalName string `uri:"name" json:"-" binding:"required" validate:"required,min=1,max=50"`
	Name         string `json:"name" validate:"required,min=1,max=50"`
	Disabled     bool   `json:"disabled"`
}

func (u UserTypeUpdateResource) GetID() int64 {
	return u.ID
}

func (u UserTypeUpdateResource) GetName() string {
	return u.Name
}

// UserTypeDeleteResource

type UserTypeDeleteResource struct {
	Name string `uri:"name" json:"-" binding:"required" validate:"required,min=1,max=50"`
}

// UserTypeResource

type UserTypeResource struct {
	Name      string    `json:"name"`
	Disabled  bool      `json:"disabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserTypeResourceBuilder

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
