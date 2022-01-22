package product

type ProductRequestFormat struct {
	Id_product_category int    `json:"id_product_category" form:"id_product_category"`
	Name                string `json:"name" form:"name"`
	Description         string `json:"description" form:"description"`
	Price               int    `json:"price" form:"price"`
	Quantity            int    `json:"quantity" form:"quantity"`
	Url_photo           string `json:"url_photo" form:"url_photo"`
}
