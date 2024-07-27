package models

type PaginationParams struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}
