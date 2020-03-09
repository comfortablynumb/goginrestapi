package user

// Structs

// Filters

type UserFindFilters struct {
	username *string
}

func (u *UserFindFilters) WithUsernamePtr(username *string) *UserFindFilters {
	u.username = username

	return u
}

func (u *UserFindFilters) WithUsername(username string) *UserFindFilters {
	return u.WithUsernamePtr(&username)
}

func (u *UserFindFilters) GetUsername() *string {
	return u.username
}

func (u *UserFindFilters) GetUsernameValue() string {
	return *u.username
}

// Options

type UserFindOptions struct {
	sortBy  *string
	sortDir *string
	offset  *int
	limit   *int
}

func (u *UserFindOptions) WithSortBy(field string, dir string) *UserFindOptions {
	return u.WithSortByPtr(&field, &dir)
}

func (u *UserFindOptions) WithSortByPtr(field *string, dir *string) *UserFindOptions {
	u.sortBy = field
	u.sortDir = dir

	return u
}

func (u *UserFindOptions) WithLimit(offset int, limit int) *UserFindOptions {
	return u.WithLimitPtr(&offset, &limit)
}

func (u *UserFindOptions) WithLimitPtr(offset *int, limit *int) *UserFindOptions {
	u.offset = offset
	u.limit = limit

	return u
}

func (u *UserFindOptions) GetSortBy() *string {
	return u.sortBy
}

func (u *UserFindOptions) GetSortByValue() string {
	return *u.sortBy
}

func (u *UserFindOptions) GetSortDir() *string {
	return u.sortDir
}

func (u *UserFindOptions) GetSortDirValue() string {
	return *u.sortDir
}

func (u *UserFindOptions) GetOffset() *int {
	return u.offset
}

func (u *UserFindOptions) GetOffsetValue() int {
	return *u.offset
}

func (u *UserFindOptions) GetLimit() *int {
	return u.limit
}

func (u *UserFindOptions) GetLimitValue() int {
	return *u.limit
}

// Static functions

func NewUserFindFilters() *UserFindFilters {
	return &UserFindFilters{}
}

func NewUserFindOptions() *UserFindOptions {
	return &UserFindOptions{}
}
