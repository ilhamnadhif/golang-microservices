package dto

import "api-gateway/model"

type UserResponse struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt model.Time `json:"created_at"`
	UpdatedAt model.Time `json:"updated_at"`
}

type UserCreateReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateReq struct {
	ID       int
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
