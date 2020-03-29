package utils

// Structs

// UserFindFilters

type UserFindFilters struct {
	username *string
}

func (u *UserFindFilters) WithUsername(username *string) *UserFindFilters {
	u.username = username

	return u
}

func (u *UserFindFilters) WithUsernameValue(username string) *UserFindFilters {
	return u.WithUsername(&username)
}

func (u *UserFindFilters) GetUsername() *string {
	return u.username
}

func (u *UserFindFilters) GetUsernameValue() string {
	return *u.username
}

// Options

// UserFindOptions

type UserFindOptions struct {
	FindOptions
}

func (f *UserFindOptions) WithSortBy(sortBy *string) *UserFindOptions {
	f.sortBy = sortBy

	return f
}

func (f *UserFindOptions) WithSortByValue(sortBy string) *UserFindOptions {
	return f.WithSortBy(&sortBy)
}

func (f *UserFindOptions) WithSortDir(sortDir *string) *UserFindOptions {
	f.sortDir = sortDir

	return f
}

func (f *UserFindOptions) WithSortDirValue(sortDir string) *UserFindOptions {
	return f.WithSortDir(&sortDir)
}

func (f *UserFindOptions) WithOffset(offset *int) *UserFindOptions {
	f.offset = offset

	return f
}

func (f *UserFindOptions) WithOffsetValue(offset int) *UserFindOptions {
	return f.WithOffset(&offset)
}

func (f *UserFindOptions) WithLimit(limit *int) *UserFindOptions {
	f.limit = limit

	return f
}

func (f *UserFindOptions) WithLimitValue(limit int) *UserFindOptions {
	return f.WithLimit(&limit)
}

func (f *UserFindOptions) WithCount(count bool) *UserFindOptions {
	f.count = count

	return f
}

func (f *UserFindOptions) GetSortBy() *string {
	return f.sortBy
}

func (f *UserFindOptions) GetSortByValue() string {
	return *f.sortBy
}

func (f *UserFindOptions) GetSortDir() *string {
	return f.sortDir
}

func (f *UserFindOptions) GetSortDirValue() string {
	return *f.sortDir
}

func (f *UserFindOptions) GetOffset() *int {
	return f.offset
}

func (f *UserFindOptions) GetOffsetValue() int {
	return *f.offset
}

func (f *UserFindOptions) GetLimit() *int {
	return f.limit
}

func (f *UserFindOptions) GetLimitValue() int {
	return *f.limit
}

func (f *UserFindOptions) IsCount() bool {
	return f.count
}

func (f *UserFindOptions) IsAsc() bool {
	return f.GetSortDir() == nil || f.GetSortDirValue() == SortDirAsc
}

func (f *UserFindOptions) IsDesc() bool {
	return !f.IsAsc()
}

// Static functions

func NewUserFindFilters() *UserFindFilters {
	return &UserFindFilters{}
}

func NewUserFindOptions() *UserFindOptions {
	return &UserFindOptions{}
}
