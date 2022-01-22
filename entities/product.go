package entities

type Products struct {
	Id                  int    `json:"id" form:"id"`
	Id_user             int    `json:"id_user" form:"id_user"`
	Id_product_category int    `json:"id_product_category" form:"id_product_category"`
	Name                string `json:"name" form:"name"`
	Description         string `json:"description" form:"description"`
	Price               int    `json:"price" form:"price"`
	Quantity            int    `json:"quantity" form:"quantity"`
	Url_photo           string `json:"url_photo" form:"url_photo"`
}

type ProductResponseFormat struct {
	Id                  int    `json:"id" form:"id"`
	Id_user             int    `json:"id_user" form:"id_user"`
	Username            string `json:"username" form:"username"`
	Id_product_category int    `json:"id_product_category" form:"id_product_category"`
	Name_category       string `json:"name_category" form:"name_category"`
	Name                string `json:"name" form:"name"`
	Description         string `json:"description" form:"description"`
	Price               int    `json:"price" form:"price"`
	Quantity            int    `json:"quantity" form:"quantity"`
	Url_photo           string `json:"url_photo" form:"url_photo"`
}
