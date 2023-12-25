package domain

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}
