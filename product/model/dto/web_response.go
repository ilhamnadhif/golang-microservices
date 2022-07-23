package dto

type PaginationReq struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}
