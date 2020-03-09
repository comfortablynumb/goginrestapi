package user

// Struct

type User struct {
	ID       int64
	Username string
	Disabled bool
}

// Static functions

func NewEmptyUser() *User {
	return &User{}
}

func NewUser(ID int64, username string, disabled bool) *User {
	return &User{
		ID:       ID,
		Username: username,
		Disabled: disabled,
	}
}
