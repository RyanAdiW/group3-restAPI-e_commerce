package order

import (
	"database/sql"
	"sirclo/groupproject/restapi/entities"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) Get(id_user int) ([]entities.CartResponseFormat, error) {

	return []entities.CartResponseFormat{}, nil
}

func (or *OrderRepository) Create(order entities.Order) error {
	//res, err := or.db.Exec("INSERT INTO order(id_user, status, total_price) VALUES(?,?,?)", order.Id_user, order.Status, order.Total_price)
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

// func (or *OrderRepository) Create(order entities.Order) error {
// 	_, err := or.db.Exec("INSERT INTO test(id_user, status, total_price) VALUES(?,?,?)", order.Id_user, order.Status, order.Total_price)
// 	return err

// }

func (or *OrderRepository) Delete(id int) error {

	return nil
}
