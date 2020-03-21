package utils

// Structs

// UserTypeFindFilters

type UserTypeFindFilters struct {
	name *string
}

func (u *UserTypeFindFilters) GetName() *string {
	return u.name
}

func (u *UserTypeFindFilters) GetNameValue() string {
	return *u.name
}

// UserTypeFindFiltersBuilder

type UserTypeFindFiltersBuilder struct {
	name *string
}

func (u *UserTypeFindFiltersBuilder) WithName(name *string) *UserTypeFindFiltersBuilder {
	u.name = name

	return u
}

func (u *UserTypeFindFiltersBuilder) WithNameValue(name string) *UserTypeFindFiltersBuilder {
	return u.WithName(&name)
}

func (u *UserTypeFindFiltersBuilder) GetName() *string {
	return u.name
}

func (u *UserTypeFindFiltersBuilder) GetNameValue() string {
	return *u.name
}

func (u *UserTypeFindFiltersBuilder) Build() *UserTypeFindFilters {
	return &UserTypeFindFilters{
		name: u.name,
	}
}

// Options

// UserTypeFindOptions

type UserTypeFindOptions struct {
	FindOptions
}

// UserTypeFindOptionsBuilder

type UserTypeFindOptionsBuilder struct {
	sortBy  *string
	sortDir *string
	offset  *int
	limit   *int
}

func (u *UserTypeFindOptionsBuilder) WithSortByValue(field string, dir string) *UserTypeFindOptionsBuilder {
	return u.WithSortBy(&field, &dir)
}

func (u *UserTypeFindOptionsBuilder) WithSortBy(field *string, dir *string) *UserTypeFindOptionsBuilder {
	u.sortBy = field
	u.sortDir = dir

	return u
}

func (u *UserTypeFindOptionsBuilder) WithLimitValue(offset int, limit int) *UserTypeFindOptionsBuilder {
	return u.WithLimit(&offset, &limit)
}

func (u *UserTypeFindOptionsBuilder) WithLimit(offset *int, limit *int) *UserTypeFindOptionsBuilder {
	u.offset = offset
	u.limit = limit

	return u
}

func (b *UserTypeFindOptionsBuilder) Build() *UserTypeFindOptions {
	return &UserTypeFindOptions{
		FindOptions{
			SortBy:  b.sortBy,
			SortDir: b.sortDir,
			Offset:  b.offset,
			Limit:   b.limit,
		},
	}
}

// Static functions

func NewUserTypeFindFiltersBuilder() *UserTypeFindFiltersBuilder {
	return &UserTypeFindFiltersBuilder{}
}

func NewUserTypeFindOptionsBuilder(defaultLimit int) *UserTypeFindOptionsBuilder {
	offset := 0

	return &UserTypeFindOptionsBuilder{
		offset: &offset,
		limit:  &defaultLimit,
	}
}
