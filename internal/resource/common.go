package resource

// Structs

type CommonFindResource struct {
	SortBy  *string `form:"sort_by"`
	SortDir *string `form:"sort_dir"`
	Offset  *int    `form:"offset"`
	Limit   *int    `form:"limit"`
}
