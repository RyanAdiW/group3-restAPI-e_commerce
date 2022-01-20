package auth

import (
	"database/sql"
	"fmt"
	_middlewares "sirclo/groupproject/restapi/delivery/middleware"
	"sirclo/groupproject/restapi/entities"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{db: db}
}

func (a *authRepository) LoginUserName(userName, password string) (string, error) {
	result, err := a.db.Query("SELECT * FROM users WHERE user_name=? AND password=?", userName, password)
	if err != nil {
		return "", err
	}
	if isExist := result.Next(); !isExist {
		return "", fmt.Errorf("id not found")
	}
	var user entities.User
	errScan := result.Scan(&user.Id, &user.Name, &user.User_name, &user.Email, &user.Password, &user.Born_date, &user.Gender)
	if errScan != nil {
		return "", errScan
	}
	token, err := _middlewares.CreateToken(user.Id, user.User_name)
	if err != nil {
		return "", err
	}
	return token, nil
}
