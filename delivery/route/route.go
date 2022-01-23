package route

import (
	"sirclo/groupproject/restapi/delivery/controller/auth"
	"sirclo/groupproject/restapi/delivery/controller/cart"
	"sirclo/groupproject/restapi/delivery/controller/product"
	"sirclo/groupproject/restapi/delivery/controller/productCategory"
	"sirclo/groupproject/restapi/delivery/controller/user"

	middlewares "sirclo/groupproject/restapi/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	loginController *auth.AuthController,
	userController *user.UserController,
	productController *product.ProductController,
	productCategoryController *productCategory.ProductCategoryController,
	cartController *cart.CartController) {

	// login
	e.POST("/login", loginController.LoginUserNameController())

	// user
	e.POST("/users", userController.CreateUserController())
	e.GET("/users/:id", userController.GetUserController(), middlewares.JWTMiddleware())
	e.PUT("/users/:id", userController.UpdateUserController(), middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userController.DeleteUserController(), middlewares.JWTMiddleware())

	// product
	e.GET("/products", productController.GetProductsController())
	e.POST("/products", productController.CreateProductController(), middlewares.JWTMiddleware())
	e.GET("/products/:id", productController.GetByIdController())
	e.PUT("/products/:id", productController.UpdateProductController(), middlewares.JWTMiddleware())
	e.DELETE("/products/:id", productController.DeleteProductController(), middlewares.JWTMiddleware())

	// cart
	e.GET("/cart", cartController.Get(), middlewares.JWTMiddleware())
	e.POST("/cart", cartController.Create(), middlewares.JWTMiddleware())
	e.PUT("/cart/:id", cartController.Update(), middlewares.JWTMiddleware())
	e.DELETE("/cart/:id", cartController.Delete(), middlewares.JWTMiddleware())

	// product category
	e.GET("/productcategory", productCategoryController.GetProductCategoryController(), middlewares.JWTMiddleware())
}
