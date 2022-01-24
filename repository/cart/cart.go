package cart

import (
	"database/sql"
	"fmt"
	"sirclo/groupproject/restapi/entities"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (cr *CartRepository) Get(id_user int) ([]entities.CartResponseFormat, error) {
	result, err := cr.db.Query(`SELECT c.id, c.id_user, c.quantity, c.total_price, p.id, pc.name_category, p.name, p.description, p.price, p.quantity, u.name , p.url_photo FROM cart as c
	INNER JOIN products as p ON c.id_product = p.id
	INNER JOIN users as u ON p.id_user = u.id
	INNER JOIN product_category as pc ON p.id_product_category = pc.id WHERE c.id_user =? AND c.id NOT IN (SELECT id_cart FROM order_detail)`, id_user)
	if err != nil {
		return nil, err
	}
	var carts []entities.CartResponseFormat
	for result.Next() {
		var cart entities.CartResponseFormat
		err := result.Scan(&cart.Id, &cart.Id_user, &cart.Quantity, &cart.Total_price, &cart.Product.Id, &cart.Product.Name_category, &cart.Product.Name, &cart.Product.Description, &cart.Product.Price, &cart.Product.Quantity, &cart.Product.Username, &cart.Product.Url_photo)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

func (cr *CartRepository) Create(cart entities.Cart) error {
	_, err := cr.db.Exec("INSERT INTO cart(id_user, id_product, quantity, total_price) VALUES(?,?,?,?)", cart.Id_user, cart.Id_product, cart.Quantity, cart.Total_price)
	return err
}

func (cr *CartRepository) Update(cart entities.Cart, id int) error {
	res, err := cr.db.Exec("UPDATE cart SET quantity=? WHERE id=?", cart.Quantity, id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}

func (cr *CartRepository) Delete(id int) error {
	res, err := cr.db.Exec("DELETE FROM cart WHERE id=?", id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}

func (cr *CartRepository) GetProductPrice(id_product int) (entities.Products, error) {
	result, err := cr.db.Query(`SELECT price FROM products WHERE id = ?`, id_product)
	if err != nil {
		return entities.Products{}, err
	}
	if isExist := result.Next(); !isExist {
		return entities.Products{}, fmt.Errorf("id not found")
	}

	var product entities.Products
	errScan := result.Scan(&product.Price)
	if errScan != nil {
		return entities.Products{}, errScan
	}
	return product, nil
}
