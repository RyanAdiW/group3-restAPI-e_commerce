package user

import (
	"database/sql"
	"fmt"
	"sirclo/groupproject/restapi/entities"
)

type userRepository struct {
	db *sql.DB
}

func NewUserReposiroty(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

// get user by id
func (ur *userRepository) GetUserById(id int) (entities.UserResponseFormat, error) {
	result, err := ur.db.Query(`SELECT id, name, username, email, born_date, gender, url_photo FROM users WHERE id = ?`, id)
	if err != nil {
		return entities.UserResponseFormat{}, err
	}
	if isExist := result.Next(); !isExist {
		return entities.UserResponseFormat{}, fmt.Errorf("id not found")
	}

	var user entities.UserResponseFormat
	errScan := result.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Born_date, &user.Gender, &user.Url_photo)
	if errScan != nil {
		return entities.UserResponseFormat{}, errScan
	}
	return user, nil
}

// insert new user
func (ur *userRepository) CreateUser(user entities.Users) error {
	_, err := ur.db.Exec("INSERT INTO users(name, username, email, password, born_date, gender, url_photo) VALUES(?,?,?,?,?,?,?)", user.Name, user.Username, user.Email, user.Password, user.Born_date, user.Gender, user.Url_photo)
	return err
}

// update user
func (ur *userRepository) UpdateUser(user entities.Users, id int) error {
	res, err := ur.db.Exec("UPDATE users SET name=?,username=?,email=?,password=?,born_date=?,gender=?, url_photo=? WHERE id=?", user.Name, user.Username, user.Email, user.Password, user.Born_date, user.Gender, user.Url_photo, id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}

// delete user
func (ur *userRepository) DeleteUser(id int) error {
	res, err := ur.db.Exec("DELETE FROM users WHERE id=?", id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}
