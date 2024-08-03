package model

type PaginationParams struct {
	Limit  int32 `form:"limit,default=50"`
	Offset int32 `form:"offset,default=0"`
}
