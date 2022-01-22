package productCategory

import (
	"database/sql"
	"sirclo/groupproject/restapi/entities"
)

type productCategoryRepository struct {
	db *sql.DB
}

func NewProductCategoryRepository(db *sql.DB) *productCategoryRepository {
	return &productCategoryRepository{db: db}
}

// get product categories
func (pcr *productCategoryRepository) GetProductCategory() ([]entities.ProductCategory, error) {
	result, err := pcr.db.Query(`SELECT id, name_category FROM product_category`)
	if err != nil {
		return nil, err
	}
	var products []entities.ProductCategory
	for result.Next() {
		var product entities.ProductCategory
		err := result.Scan(&product.Id, &product.Name_category)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
