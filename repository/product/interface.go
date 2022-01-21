package product

import (
	"sirclo/groupproject/restapi/entities"
)

type UserRepository interface {
	GetProducts() ([]entities.ProductResponseFormat, error)
	GetProductById(id int) (entities.ProductResponseFormat, error)
	CreateProduct(user entities.Products) error
	UpdateProduct(user entities.Products, id int) error
	DeleteProduct(id int) error
}
