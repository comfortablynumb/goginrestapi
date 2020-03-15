package utils

// Structs

// Filters

type UserTypeFindFilters struct {
	name *string
}

func (u *UserTypeFindFilters) WithNamePtr(name *string) *UserTypeFindFilters {
	u.name = name

	return u
}

func (u *UserTypeFindFilters) WithName(name string) *UserTypeFindFilters {
	return u.WithNamePtr(&name)
}

func (u *UserTypeFindFilters) GetName() *string {
	return u.name
}

func (u *UserTypeFindFilters) GetNameValue() string {
	return *u.name
}

// Options

type UserTypeFindOptions struct {
	sortBy  *string
	sortDir *string
	offset  *int
	limit   *int
}

func (u *UserTypeFindOptions) WithSortBy(field string, dir string) *UserTypeFindOptions {
	return u.WithSortByPtr(&field, &dir)
}

func (u *UserTypeFindOptions) WithSortByPtr(field *string, dir *string) *UserTypeFindOptions {
	u.sortBy = field
	u.sortDir = dir

	return u
}

func (u *UserTypeFindOptions) WithLimit(offset int, limit int) *UserTypeFindOptions {
	return u.WithLimitPtr(&offset, &limit)
}

func (u *UserTypeFindOptions) WithLimitPtr(offset *int, limit *int) *UserTypeFindOptions {
	u.offset = offset
	u.limit = limit

	return u
}

func (u *UserTypeFindOptions) GetSortBy() *string {
	return u.sortBy
}

func (u *UserTypeFindOptions) GetSortByValue() string {
	return *u.sortBy
}

func (u *UserTypeFindOptions) GetSortDir() *string {
	return u.sortDir
}

func (u *UserTypeFindOptions) GetSortDirValue() string {
	return *u.sortDir
}

func (u *UserTypeFindOptions) GetOffset() *int {
	return u.offset
}

func (u *UserTypeFindOptions) GetOffsetValue() int {
	return *u.offset
}

func (u *UserTypeFindOptions) GetLimit() *int {
	return u.limit
}

func (u *UserTypeFindOptions) GetLimitValue() int {
	return *u.limit
}

// Static functions

func NewUserTypeFindFilters() *UserTypeFindFilters {
	return &UserTypeFindFilters{}
}

func NewUserTypeFindOptions() *UserTypeFindOptions {
	return &UserTypeFindOptions{}
}
