package models

import "github.com/golang-jwt/jwt/v5"

type ClaimsTokenData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}
