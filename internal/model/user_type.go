package model

import "time"

// Interfaces

type EntityWithUserType interface {
	SetUserType(userType UserType)
}

// Structs

type UserType struct {
	ID        int64
	Name      string
	Disabled  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserTypeBuilder struct {
	id        int64
	name      string
	disabled  bool
	createdAt time.Time
	updatedAt time.Time
}

func (b *UserTypeBuilder) WithID(ID int64) *UserTypeBuilder {
	b.id = ID

	return b
}

func (b *UserTypeBuilder) WithName(name string) *UserTypeBuilder {
	b.name = name

	return b
}

func (b *UserTypeBuilder) WithDisabled(disabled bool) *UserTypeBuilder {
	b.disabled = disabled

	return b
}

func (b *UserTypeBuilder) WithCreatedAt(createdAt time.Time) *UserTypeBuilder {
	b.createdAt = createdAt

	return b
}

func (b *UserTypeBuilder) WithUpdatedAt(updatedAt time.Time) *UserTypeBuilder {
	b.updatedAt = updatedAt

	return b
}

func (b *UserTypeBuilder) Build() *UserType {
	return &UserType{
		ID:        b.id,
		Name:      b.name,
		Disabled:  b.disabled,
		CreatedAt: b.createdAt,
		UpdatedAt: b.updatedAt,
	}
}

// Static functions

func NewUserTypeBuilder() *UserTypeBuilder {
	return &UserTypeBuilder{}
}
