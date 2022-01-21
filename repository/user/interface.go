package user

import (
	"sirclo/groupproject/restapi/entities"
)

type UserRepository interface {
	GetUserById(id int) (entities.UserResponseFormat, error)
	CreateUser(user entities.Users) error
	UpdateUser(user entities.Users, id int) error
	DeleteUser(id int) error
}
