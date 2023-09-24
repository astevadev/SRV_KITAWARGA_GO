package main

import (
	"kitawarga/cmd/config"
	"kitawarga/cmd/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	config.InitDB()

	e.Use(controller.LoggingRequest)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controller.Home)

	e.POST("/login", controller.Login)

	r := e.Group("/api")

	r.Use(echojwt.WithConfig(config.ConfigRestricted))
	r.POST("/", controller.Restricted)
	r.POST("/getAllUsers", controller.GetUsers)
	r.POST("/addUser", controller.CreateUser)
	e.Logger.Fatal(e.Start(":8080"))
}
