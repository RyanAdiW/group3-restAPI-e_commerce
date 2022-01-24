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
	result, err := a.db.Query("SELECT * FROM users WHERE username=? AND password=?", userName, password)
	if err != nil {
		return "", err
	}
	if isExist := result.Next(); !isExist {
		return "", fmt.Errorf("id not found")
	}
	var user entities.Users
	errScan := result.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Born_date, &user.Gender, &user.Url_photo)
	if errScan != nil {
		return "", errScan
	}
	token, err := _middlewares.CreateToken(user.Id, user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authRepository) GetPasswordByUsername(userName string) (string, error) {
	result, err := a.db.Query("SELECT * FROM users WHERE username=?", userName)
	if err != nil {
		return "", err
	}
	if isExist := result.Next(); !isExist {
		return "", fmt.Errorf("id not found")
	}
	var user entities.Users
	errScan := result.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Born_date, &user.Gender, &user.Url_photo)
	if errScan != nil {
		return "", errScan
	}
	password := user.Password
	return password, nil
}

func (a *authRepository) GetIdByUsername(userName string) (int, error) {
	result, err := a.db.Query("SELECT id FROM users WHERE username=?", userName)
	if err != nil {
		return 0, err
	}
	if isExist := result.Next(); !isExist {
		return 0, fmt.Errorf("id not found")
	}
	var user entities.Users
	errScan := result.Scan(&user.Id)
	if errScan != nil {
		return 0, errScan
	}
	userId := user.Id
	return userId, nil
}
