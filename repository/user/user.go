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
	result, err := ur.db.Query(`SELECT id, name, user_name, email, born_date, gender FROM users WHERE id = ?`, id)
	if err != nil {
		return entities.UserResponseFormat{}, err
	}
	if isExist := result.Next(); !isExist {
		return entities.UserResponseFormat{}, fmt.Errorf("id not found")
	}

	var user entities.UserResponseFormat
	errScan := result.Scan(&user.Id, &user.Name, &user.User_name, &user.Email, &user.Born_date, &user.Gender)
	if errScan != nil {
		return entities.UserResponseFormat{}, errScan
	}
	return user, nil
}

// insert new user
func (ur *userRepository) CreateUser(user entities.User) error {
	_, err := ur.db.Exec("INSERT INTO users(name, user_name, email, password, born_date, gender) VALUES(?,?,?,?,?)", user.Name, user.User_name, user.Email, user.Password, user.Born_date, user.Gender)
	return err
}

// update user
func (ur *userRepository) UpdateUser(user entities.User, id int) error {
	res, err := ur.db.Exec("UPDATE users SET name=?,user_name=?,email=?,password=?,born_date=?,gender=? WHERE id=?", user.Name, user.User_name, user.Email, user.Password, user.Born_date, user.Gender, id)
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

func main() {

}