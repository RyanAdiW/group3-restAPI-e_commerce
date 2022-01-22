package entities

type Cart struct {
	Id         int `json:"id" form:"id"`
	Id_user    int `json:"id_user" form:"id_user"`
	Id_product int `json:"id_product" form:"id_product"`
	Quantity   int `json:"quantity" form:"quantity"`
}

type CartResponseFormat struct {
	Id       int `json:"id" form:"id"`
	Id_user  int `json:"id_user" form:"id_user"`
	Product  Products
	Quantity int `json:"quantity" form:"quantity"`
}

type Products struct {
	Id               int    `json:"id" form:"id"`
	User_owner       string `json:"user_owner" form:"user_owner"`
	Product_category string `json:"id_product_category" form:"id_product_category"`
	Name             string `json:"name" form:"name"`
	Price            int    `json:"price" form:"price"`
	Quantity         int    `json:"quantity" form:"quantity"`
	Description      string `json:"description" form:"description"`
	Url_photo        string `json:"url_photo" form:"url_photo"`
}
