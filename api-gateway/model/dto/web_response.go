package dto

import "net/http"

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"message"`
	Data   interface{} `json:"data"`
}

type Metadata struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	RowCount int64 `json:"row_count"`
}

type PaginationReq struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}

type ListAllResponse struct {
	List     interface{} `json:"list"`
	Metadata Metadata    `json:"metadata"`
}

func WebResponseSuccess(data interface{}) WebResponse {
	return WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   data,
	}
}
