package auth

type LoginUserNameRequestFormat struct {
	User_name string `json:"user_name" form:"user_name"`
	Password  string `json:"password" form:"password"`
}
