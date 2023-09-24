package main

import (
	"kitawarga/cmd/handlers"
	"kitawarga/cmd/storage"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	storage.InitDB()

	e.Use(handlers.LoggingRequest)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.Home)

	e.POST("/login", handlers.Login)

	r := e.Group("/api")

	r.Use(echojwt.WithConfig(handlers.ConfigRestricted))
	r.POST("/", handlers.Restricted)
	r.POST("/getAllUsers", handlers.GetUsers)
	r.POST("/addUser", handlers.CreateUser)
	e.Logger.Fatal(e.Start(":8080"))
}
