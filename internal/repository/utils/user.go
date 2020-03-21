package utils

// Structs

// UserFindFilters

type UserFindFilters struct {
	username *string
}

func (u *UserFindFilters) GetUsername() *string {
	return u.username
}

func (u *UserFindFilters) GetUsernameValue() string {
	return *u.username
}

// UserFindFiltersBuilder

type UserFindFiltersBuilder struct {
	username *string
}

func (u *UserFindFiltersBuilder) WithUsername(username *string) *UserFindFiltersBuilder {
	u.username = username

	return u
}

func (u *UserFindFiltersBuilder) WithUsernameValue(username string) *UserFindFiltersBuilder {
	return u.WithUsername(&username)
}

func (u *UserFindFiltersBuilder) GetUsername() *string {
	return u.username
}

func (u *UserFindFiltersBuilder) GetUsernameValue() string {
	return *u.username
}

func (u *UserFindFiltersBuilder) Build() *UserFindFilters {
	return &UserFindFilters{
		username: u.username,
	}
}

// Options

// UserFindOptions

type UserFindOptions struct {
	FindOptions
}

// UserFindOptionsBuilder

type UserFindOptionsBuilder struct {
	sortBy  *string
	sortDir *string
	offset  *int
	limit   *int
}

func (u *UserFindOptionsBuilder) WithSortByValue(field string, dir string) *UserFindOptionsBuilder {
	return u.WithSortBy(&field, &dir)
}

func (u *UserFindOptionsBuilder) WithSortBy(field *string, dir *string) *UserFindOptionsBuilder {
	u.sortBy = field
	u.sortDir = dir

	return u
}

func (u *UserFindOptionsBuilder) WithLimitValue(offset int, limit int) *UserFindOptionsBuilder {
	return u.WithLimit(&offset, &limit)
}

func (u *UserFindOptionsBuilder) WithLimit(offset *int, limit *int) *UserFindOptionsBuilder {
	u.offset = offset
	u.limit = limit

	return u
}

func (b *UserFindOptionsBuilder) Build() *UserFindOptions {
	return &UserFindOptions{
		FindOptions{
			SortBy:  b.sortBy,
			SortDir: b.sortDir,
			Offset:  b.offset,
			Limit:   b.limit,
		},
	}
}

// Static functions

func NewUserFindFiltersBuilder() *UserFindFiltersBuilder {
	return &UserFindFiltersBuilder{}
}

func NewUserFindOptionsBuilder(defaultLimit int) *UserFindOptionsBuilder {
	offset := 0

	return &UserFindOptionsBuilder{
		offset: &offset,
		limit:  &defaultLimit,
	}
}
