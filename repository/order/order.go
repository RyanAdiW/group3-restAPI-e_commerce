package order

import (
	"database/sql"
	"fmt"
	"sirclo/groupproject/restapi/entities"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) Get(id_user int) ([]entities.OrderResponseFormat, error) {

	result, err := or.db.Query(`SELECT c.id, c.id_user, c.quantity, c.total_price, p.id, pc.name_category, p.name, p.description, p.price, p.quantity, u.name , p.url_photo, o.id, o.status, od.order_date FROM cart as c
	INNER JOIN products as p ON c.id_product = p.id
	INNER JOIN users as u ON p.id_user = u.id
	INNER JOIN product_category as pc ON p.id_product_category = pc.id
	INNER JOIN order_detail as od ON c.id = od.id_cart
	INNER JOIN orders as o ON od.id_order = o.id WHERE c.id_user =? AND c.id IN (SELECT id_cart FROM order_detail as od INNER JOIN orders as o ON od.id_order = o.id WHERE o.status = "DONE" OR o.status = "CANCELLED")`, id_user)
	if err != nil {
		return nil, err
	}
	var orders []entities.OrderResponseFormat
	for result.Next() {
		var order entities.OrderResponseFormat
		err := result.Scan(&order.Id, &order.Id_user, &order.Quantity, &order.Total_price, &order.Product.Id, &order.Product.Name_category, &order.Product.Name, &order.Product.Description, &order.Product.Price, &order.Product.Quantity, &order.Product.Username, &order.Product.Url_photo, &order.Id_order, &order.Status, &order.Order_date)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (or *OrderRepository) Create(order entities.Order) error {
	res, err := or.db.Exec("INSERT INTO orders(id_user, status, total_price) VALUES(?,?,?)", order.Id_user, order.Status, order.Total_price)
	if err != nil {
		return err
	}
	orderId, _ := res.LastInsertId()
	_, err = or.db.Exec("INSERT INTO address_delivery(id, street, city, state, zip) VALUES(?,?,?,?,?)", orderId, order.Address_delivery.Street, order.Address_delivery.City, order.Address_delivery.State, order.Address_delivery.Zip)
	if err != nil {
		return err
	}
	_, err = or.db.Exec("INSERT INTO credit_cards(id, type, name, number, cvv, month, year) VALUES(?,?,?,?,?,?,?)", orderId, order.Credit_card.Type, order.Credit_card.Name, order.Credit_card.Number, order.Credit_card.Cvv, order.Credit_card.Month, order.Credit_card.Year)

	for i := 0; i < len(order.Id_cart); i++ {
		_, err = or.db.Exec("INSERT INTO order_detail(id_order, id_cart) VALUES(?,?)", orderId, order.Id_cart[i])
		if err != nil {
			return err
		}
	}

	return err
}

func (or *OrderRepository) Update(order entities.Order, id int) error {
	res, err := or.db.Exec("UPDATE orders SET status=? WHERE id=?", order.Status, id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Errorf("id not found")
	}
	return err
}
