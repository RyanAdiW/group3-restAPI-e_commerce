package product

import (
	"database/sql"
	"fmt"
	"sirclo/groupproject/restapi/entities"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepositpry(db *sql.DB) *productRepository {
	return &productRepository{db: db}
}

// 1. get all products
func (pr *productRepository) GetProducts() ([]entities.ProductResponseFormat, error) {
	result, err := pr.db.Query(`
	SELECT p.id as id_product, u.id as id_user, u.username, pc.id as id_product_category, 
	pc.name_category, p.name, p.description, p.price, p.quantity, p.url_photo
	FROM products p 
	INNER JOIN users as u ON (p.id_user = u.id)
	INNER JOIN product_category as pc ON (p.id_product_category = pc.id)`)

	if err != nil {
		return nil, err
	}
	var products []entities.ProductResponseFormat
	for result.Next() {
		var product entities.ProductResponseFormat
		err := result.Scan(&product.Id, &product.Id_user, &product.Username, &product.Id_product_category, &product.Name_category, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.Url_photo)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// 2. get product by id
func (pr *productRepository) GetProductById(id int) (entities.ProductResponseFormat, error) {
	result, err := pr.db.Query(`
	SELECT p.id as id_product, u.id as id_user, u.username, pc.id as id_product_category, 
	pc.name_category, p.name, p.description, p.price, p.quantity, p.url_photo
	FROM products p 
	INNER JOIN users as u ON (p.id_user = u.id)
	INNER JOIN product_category as pc ON (p.id_product_category = pc.id)
	WHERE p.id=?`, id)
	if err != nil {
		return entities.ProductResponseFormat{}, err
	}

	if isExist := result.Next(); !isExist {
		return entities.ProductResponseFormat{}, fmt.Errorf("id not found")
	}

	var product entities.ProductResponseFormat
	errScan := result.Scan(&product.Id, &product.Id_user, &product.Username, &product.Id_product_category, &product.Name_category, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.Url_photo)
	if err != nil {
		return entities.ProductResponseFormat{}, errScan
	}
	return product, nil
}

// 3. create product
func (pr *productRepository) CreateProduct(product entities.Products) error {
	_, err := pr.db.Exec(`INSERT INTO products(id_user, id_product_category, name, description, price, quantity, url_photo)
	VALUES(?,?,?,?,?,?,?)`, product.Id_user, product.Id_product_category, product.Name, product.Description, product.Price, product.Quantity, product.Url_photo)
	return err
}

// 4. update product
func (pr *productRepository) UpdateProduct(product entities.Products, id int) error {
	res, err := pr.db.Exec(`UPDATE products SET id_product_category=?, name=?, description=?, price=?, quantity=?, url_photo=? WHERE id=?`, product.Id_product_category, product.Name, product.Description, product.Price, product.Quantity, product.Url_photo, id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}

// 5. Delete product
func (pr *productRepository) DeleteProduct(id int) error {
	res, err := pr.db.Exec("DELETE FROM products WHERE id=?", id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}
