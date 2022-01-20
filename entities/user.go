package entities

type User struct {
	Id        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	User_name string `json:"user_name" form:"user_name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Born_date string `json:"born_date" form:"born_date"`
	Gender    string `json:"gender" form:"gender"`
}

type UserResponseFormat struct {
	Id        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	User_name string `json:"user_name" form:"user_name"`
	Email     string `json:"email" form:"email"`
	Born_date string `json:"born_date" form:"born_date"`
	Gender    string `json:"gender" form:"gender"`
}
