package product

import (
	"sirclo/groupproject/restapi/entities"
)

type ProductRepository interface {
	GetProducts() ([]entities.ProductResponseFormat, error)
	GetProductById(id int) (entities.ProductResponseFormat, error)
	CreateProduct(product entities.Products) error
	UpdateProduct(product entities.Products, id int) error
	DeleteProduct(id int) error
}
