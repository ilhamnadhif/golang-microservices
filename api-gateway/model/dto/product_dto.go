package dto

import "api-gateway/model"

type ProductResponse struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       int        `json:"price"`
	CreatedAt   model.Time `json:"created_at"`
	UpdatedAt   model.Time `json:"updated_at"`
}

type ProductCreateReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ProductUpdateReq struct {
	ID          int
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
