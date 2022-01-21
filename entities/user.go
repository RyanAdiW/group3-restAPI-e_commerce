package entities

type User struct {
	Id         int    `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Username   string `json:"username" form:"username"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Birth_date string `json:"birth_date" form:"birth_date"`
	Gender     string `json:"gender" form:"gender"`
	Url_photo  string `json:"url_photo" form:"url_photo"`
}

type UserResponseFormat struct {
	Id         int    `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Username   string `json:"username" form:"username"`
	Email      string `json:"email" form:"email"`
	Birth_date string `json:"birth_date" form:"birth_date"`
	Gender     string `json:"gender" form:"gender"`
	Url_photo  string `json:"url_photo" form:"url_photo"`
}
