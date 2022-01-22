package route

import (
	"sirclo/groupproject/restapi/delivery/controller/auth"
	"sirclo/groupproject/restapi/delivery/controller/product"
	"sirclo/groupproject/restapi/delivery/controller/user"

	middlewares "sirclo/groupproject/restapi/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	loginController *auth.AuthController,
	userController *user.UserController,
	productController *product.ProductController) {

	// login
	e.POST("/login", loginController.LoginUserNameController())

	// user
	e.POST("/users", userController.CreateUserController())
	e.GET("/users/:id", userController.GetUserController(), middlewares.JWTMiddleware())
	e.PUT("/users/:id", userController.UpdateUserController(), middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userController.DeleteUserController(), middlewares.JWTMiddleware())

	// product
	e.GET("/users", productController.GetProductsController())
	e.POST("/products", productController.CreateProductController(), middlewares.JWTMiddleware())
	e.GET("/users/:id", productController.GetByIdController())
	e.PUT("/users/:id", productController.UpdateProductController(), middlewares.JWTMiddleware())
	e.DELETE("/users/:id", productController.DeleteProductController(), middlewares.JWTMiddleware())
}
