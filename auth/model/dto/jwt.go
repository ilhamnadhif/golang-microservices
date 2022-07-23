package dto

import "github.com/golang-jwt/jwt"

type JWTCustomClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
