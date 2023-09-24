package controller

import (
	"kitawarga/cmd/config"
	"kitawarga/cmd/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

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
