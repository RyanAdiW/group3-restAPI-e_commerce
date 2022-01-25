package cart

import (
	"sirclo/groupproject/restapi/entities"
)

type Cart interface {
	Get(id_user int) ([]entities.CartResponseFormat, error)
	Create(cart entities.Cart) error
	Update(cart entities.Cart, id int) error
	Delete(id int) error
	GetProductPrice(id_product int) (entities.Products, error)
	GetProductFromCart(id_user, id_product int) (entities.Cart, string, error)
}
