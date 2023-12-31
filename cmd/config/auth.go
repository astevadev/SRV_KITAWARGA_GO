package config

import (
	"kitawarga/cmd/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var ConfigRestricted = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(models.ClaimsTokenData)
	},
	SigningKey: []byte(KeySecret),
}
