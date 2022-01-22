package productCategory

import (
	"sirclo/groupproject/restapi/entities"
)

type ProductCategoryRepository interface {
	GetProductCategory() ([]entities.ProductCategory, error)
}
