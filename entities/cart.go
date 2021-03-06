package entities

type Cart struct {
	Id          int `json:"id" form:"id"`
	Id_user     int `json:"id_user" form:"id_user"`
	Id_product  int `json:"id_product" form:"id_product"`
	Quantity    int `json:"quantity" form:"quantity"`
	Total_price int `json:"total_price" form:"total_price"`
}

type CartResponseFormat struct {
	Id          int `json:"id" form:"id"`
	Id_user     int `json:"id_user" form:"id_user"`
	Product     ProductResponseFormat
	Quantity    int `json:"quantity" form:"quantity"`
	Total_price int `json:"total_price" form:"total_price"`
}
