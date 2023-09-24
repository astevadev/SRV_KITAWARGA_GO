package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedDate string `json:"createddate"`
	UpdatedDate string `json:"updateddate"`
}

type Users struct {
	Users []User `json:"users"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}
