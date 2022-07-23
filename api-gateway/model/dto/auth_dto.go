package dto

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
