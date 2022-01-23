package cart

type CartRequestFormat struct {
	Id_product int `json:"id_product" form:"id_product"`
	Quantity   int `json:"quantity" form:"quantity"`
}
