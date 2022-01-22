package cart

import (
	"database/sql"
	"fmt"
	"sirclo/groupproject/restapi/entities"
)

type cartRepository struct {
	db *sql.DB
}

func NewCartReposiroty(db *sql.DB) *cartRepository {
	return &cartRepository{db: db}
}

func (cr *cartRepository) Get(id_user int) ([]entities.CartResponseFormat, error) {
	result, err := cr.db.Query(`SELECT c.id, c.id_user, c.quantity, p.id, p.product_category, p.name, p.description, p.price, p.quantity, u.name , p.url_photo FROM cart as c
	INNER JOIN products as p ON c.id_product = p.id
	INNER JOIN users as u ON p.id_user = u.id WHERE c.id_user =?`, id_user)
	if err != nil {
		return nil, err
	}
	var carts []entities.CartResponseFormat
	for result.Next() {
		var cart entities.CartResponseFormat
		err := result.Scan(&cart.Id, &cart.Id_user, &cart.Quantity, &cart.Product.Id, &cart.Product.Product_category, &cart.Product.Name, &cart.Product.Description, &cart.Product.Price, &cart.Product.Quantity, &cart.Product.User_owner, &cart.Product.Url_photo)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

func (cr *cartRepository) Create(cart entities.Cart) error {
	_, err := cr.db.Exec("INSERT INTO cart(id_user, id_product, quantity) VALUES(?,?,?)", cart.Id_user, cart.Id_product, cart.Quantity)
	return err
}

func (cr *cartRepository) Update(cart entities.Cart, id int) error {
	res, err := cr.db.Exec("UPDATE cart SET quantity=? WHERE id=?", cart.Quantity, id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}

func (cr *cartRepository) Delete(id int) error {
	res, err := cr.db.Exec("DELETE FROM cart WHERE id=?", id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}
