package user

import (
	"sirclo/groupproject/restapi/entities"
)

type UserRepository interface {
	GetUserById(id int) (entities.UserResponseFormat, error)
	CreateUser(user entities.User) error
	UpdateUser(user entities.User, id int) error
	DeleteUser(id int) error
}
