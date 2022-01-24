package order

import (
	"sirclo/groupproject/restapi/entities"
)

type Order interface {
	Get(id_user int) ([]entities.CartResponseFormat, error)
	Create(order entities.Order) error
	Delete(id int) error
}
