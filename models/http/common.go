package http

import (
	"github.com/golang-jwt/jwt/v4"
)

type AuthClaim struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Bd       string `json:"bd"`
	TokenKey string `json:"token_key"`
	jwt.RegisteredClaims
}

type Page struct {
	Now    string
	Domain string
}
