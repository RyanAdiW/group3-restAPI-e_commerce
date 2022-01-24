package order

import (
	"sirclo/groupproject/restapi/entities"
)

type OrderRequestFormat struct {
	Id_cart          []int                     `json:"id_cart" form:"id_cart"`
	Total_price      int                       `json:"total_price" form:"total_price"`
	Address_delivery entities.Address_delivery `json:"address_delivery" form:"address_delivery"`
	Credit_card      entities.Credit_card      `json:"credit_card" form:"credit_card"`
}
