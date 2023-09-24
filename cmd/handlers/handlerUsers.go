package handlers

import (
	"fmt"
	"kitawarga/cmd/models"
	"kitawarga/cmd/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := models.Users{}
	c.Bind(&users)
	rowUser, err := repositories.GetUsers(users)
	if err != nil {
		return c.JSON(http.StatusNoContent, err.Error())
	}
	fmt.Println(rowUser)
	return c.JSON(http.StatusOK, rowUser)
}

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	newUser, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(newUser)
	return c.JSON(http.StatusCreated, newUser)
}
