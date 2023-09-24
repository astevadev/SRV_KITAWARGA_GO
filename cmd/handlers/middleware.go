package handlers

import (
	"fmt"
	"kitawarga/cmd/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var LoggingRequest = middleware.LoggerWithConfig(middleware.LoggerConfig{
	Format: "method=${method}, uri=${uri}, status=${status}\n",
})

func LogRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		stop := time.Now()
		fmt.Printf("Method: %s,\t URI= %s,\t Status= %s\n", c.Request().Method, c.Request().URL, c.Response().Status, stop.Sub(start))
		return err
	}
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.ClaimsTokenData)
	name := claims.Username
	fmt.Println("Hallo Dimas, AKU UDAH BISA GOLANG !!!", name)
	return c.String(http.StatusOK, "Hallo Dimas, AKU UDAH BISA GOLANG !!!!")
}
