package entities

type ProductCategory struct {
	Id            int    `json:"id" form:"id"`
	Name_category string `json:"name_category" form:"name_category"`
}
