package entities

type Order struct {
	Id               int              `json:"id" form:"id"`
	Id_user          int              `json:"id_user" form:"id_user"`
	Status           string           `json:"status" form:"status"`
	Total_price      int              `json:"total_price" form:"total_price"`
	Address_delivery Address_delivery `json:"address_delivery" form:"address_delivery"`
	Credit_card      Credit_card      `json:"credit_card" form:"credit_card"`
	Id_cart          []int            `json:"id_cart" form:"id_cart"`
}

type OrderResponseFormat struct {
	Id          int `json:"id" form:"id"`
	Id_user     int `json:"id_user" form:"id_user"`
	Product     ProductResponseFormat
	Quantity    int    `json:"quantity" form:"quantity"`
	Total_price int    `json:"total_price" form:"total_price"`
	Id_order    int    `json:"id_order" form:"id_order"`
	Status      string `json:"status" form:"status"`
	Order_date  string `json:"order_date" form:"order_date"`
}

type Address_delivery struct {
	Street string `json:"street" form:"street"`
	City   string `json:"city" form:"city"`
	State  string `json:"state" form:"state"`
	Zip    int    `json:"zip" form:"zip"`
}

type Credit_card struct {
	Type   string `json:"type" form:"type"`
	Name   string `json:"name" form:"name"`
	Number string `json:"number" form:"number"`
	Cvv    int    `json:"cvv" form:"cvv"`
	Month  int    `json:"month" form:"month"`
	Year   int    `json:"year" form:"year"`
}
