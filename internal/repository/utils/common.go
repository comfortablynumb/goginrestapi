package utils

// Constants

const (
	SortDirAsc  = "ASC"
	SortDirDesc = "DESC"
)

// Structs

// FindOptions

type FindOptions struct {
	sortBy  *string
	sortDir *string
	offset  *int
	limit   *int
	count   bool
}
