package repositories

import (
	"database/sql"
	"kitawarga/cmd/config"
	"kitawarga/cmd/models"
	"time"
)

var rows *sql.Rows

func GetUsers(users models.Users) (models.Users, error) {
	sqlStatement := "SELECT id, name, email, password, createddate, updateddate FROM USERS"
	rows, err := config.DB.Query(sqlStatement)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	result := models.Users{}
	for rows.Next() {
		user := models.User{}
		err2 := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedDate, &user.UpdatedDate)
		if err2 != nil {
			return users, err
		}
		result.Users = append(result.Users, user)
	}
	return result, nil
}

func CreateUser(user models.User) (models.User, error) {
	sqlStatement := "INSERT INTO USERS (name, email, password, createddate) VALUES ($1, $2, $3, $4) RETURNING id"
	err := config.DB.QueryRow(sqlStatement, user.Name, user.Email, user.Password, time.Now()).Scan(&user.Id)
	if err != nil {
		return user, err
	}
	user.CreatedDate = time.Now().Format("2006-01-02 15:04:05")
	return user, nil
}
