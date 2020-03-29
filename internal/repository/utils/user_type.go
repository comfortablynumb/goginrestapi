package utils

import "strings"

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

func (u *UserTypeFindFilters) WithName(name *string) *UserTypeFindFilters {
	u.name = name

	return u
}

func (u *UserTypeFindFilters) WithNameValue(name string) *UserTypeFindFilters {
	return u.WithName(&name)
}

// Options

// UserTypeFindOptions

type UserTypeFindOptions struct {
	FindOptions
}

func (f *UserTypeFindOptions) WithSortBy(sortBy *string) *UserTypeFindOptions {
	f.sortBy = sortBy

	return f
}

func (f *UserTypeFindOptions) WithSortByValue(sortBy string) *UserTypeFindOptions {
	return f.WithSortBy(&sortBy)
}

func (f *UserTypeFindOptions) WithSortDir(sortDir *string) *UserTypeFindOptions {
	if sortDir != nil {
		*sortDir = strings.ToUpper(*sortDir)
	}

	f.sortDir = sortDir

	return f
}

func (f *UserTypeFindOptions) WithSortDirValue(sortDir string) *UserTypeFindOptions {
	return f.WithSortDir(&sortDir)
}

func (f *UserTypeFindOptions) WithOffset(offset *int) *UserTypeFindOptions {
	f.offset = offset

	return f
}

func (f *UserTypeFindOptions) WithOffsetValue(offset int) *UserTypeFindOptions {
	return f.WithOffset(&offset)
}

func (f *UserTypeFindOptions) WithLimit(limit *int) *UserTypeFindOptions {
	f.limit = limit

	return f
}

func (f *UserTypeFindOptions) WithLimitValue(limit int) *UserTypeFindOptions {
	return f.WithLimit(&limit)
}

func (f *UserTypeFindOptions) WithCount(count bool) *UserTypeFindOptions {
	f.count = count

	return f
}

func (f *UserTypeFindOptions) GetSortBy() *string {
	return f.sortBy
}

func (f *UserTypeFindOptions) GetSortByValue() string {
	return *f.sortBy
}

func (f *UserTypeFindOptions) GetSortDir() *string {
	return f.sortDir
}

func (f *UserTypeFindOptions) GetSortDirValue() string {
	return *f.sortDir
}

func (f *UserTypeFindOptions) GetOffset() *int {
	return f.offset
}

func (f *UserTypeFindOptions) GetOffsetValue() int {
	return *f.offset
}

func (f *UserTypeFindOptions) GetLimit() *int {
	return f.limit
}

func (f *UserTypeFindOptions) GetLimitValue() int {
	return *f.limit
}

func (f *UserTypeFindOptions) IsCount() bool {
	return f.count
}

func (f *UserTypeFindOptions) IsAsc() bool {
	return f.GetSortDir() != nil && f.GetSortDirValue() == SortDirAsc
}

func (f *UserTypeFindOptions) IsDesc() bool {
	return f.GetSortDir() != nil && f.GetSortDirValue() == SortDirDesc
}

// Static functions

func NewUserTypeFindFilters() *UserTypeFindFilters {
	return &UserTypeFindFilters{}
}

func NewUserTypeFindOptions() *UserTypeFindOptions {
	return &UserTypeFindOptions{}
}
