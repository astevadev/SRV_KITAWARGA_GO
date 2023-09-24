package handlers

import (
	"kitawarga/cmd/config"
	"kitawarga/cmd/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var ConfigRestricted = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(models.ClaimsTokenData)
	},
	SigningKey: []byte(config.KeySecret),
}

func Login(c echo.Context) error {
	logindata := models.ClaimsTokenData{}
	c.Bind(&logindata)
	// Check in your db if the user exists or not
	if logindata.Username == "jon" && logindata.Password == "password" {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = "jon"
		claims["password"] = "password"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(config.KeySecret))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"claims_token": t,
		})
	}
	return echo.ErrUnauthorized
}
