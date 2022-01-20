package route

import (
	"sirclo/groupproject/restapi/delivery/controller/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	userController *user.UserController) {

	// user
	e.POST("/users", userController.CreateUserController())
	e.GET("/users/:id", userController.GetUserController())
	e.PUT("/users/:id", userController.UpdateUserController())
	e.DELETE("/users/:id", userController.DeleteUserController())

}
