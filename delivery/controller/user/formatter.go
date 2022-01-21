package user

type UserRequestFormat struct {
	Name       string `json:"name" form:"name"`
	Username   string `json:"username" form:"username"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Birth_date string `json:"birth_date" form:"birth_date"`
	Gender     string `json:"gender" form:"gender"`
}
