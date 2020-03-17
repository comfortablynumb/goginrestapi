package model

import "time"

// Structs

type User struct {
	ID        int64
	Username  string
	UserType  UserType
	Disabled  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) SetUserType(userType UserType) {

}

type UserBuilder struct {
	id        int64
	username  string
	userType  UserType
	disabled  bool
	createdAt time.Time
	updatedAt time.Time
}

func (b *UserBuilder) WithID(ID int64) *UserBuilder {
	b.id = ID

	return b
}

func (b *UserBuilder) WithUsername(username string) *UserBuilder {
	b.username = username

	return b
}

func (b *UserBuilder) WithUserType(userType UserType) *UserBuilder {
	b.userType = userType

	return b
}

func (b *UserBuilder) WithDisabled(disabled bool) *UserBuilder {
	b.disabled = disabled

	return b
}

func (b *UserBuilder) WithCreatedAt(createdAt time.Time) *UserBuilder {
	b.createdAt = createdAt

	return b
}

func (b *UserBuilder) WithUpdatedAt(updatedAt time.Time) *UserBuilder {
	b.updatedAt = updatedAt

	return b
}

func (b *UserBuilder) Build() *User {
	return &User{
		ID:        b.id,
		Username:  b.username,
		UserType:  b.userType,
		Disabled:  b.disabled,
		CreatedAt: b.createdAt,
		UpdatedAt: b.updatedAt,
	}
}

// Static functions

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}
