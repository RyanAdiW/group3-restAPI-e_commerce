package main

import (
	"os"
	"sirclo/groupproject/restapi/config"
	"sirclo/groupproject/restapi/delivery/route"

	_authController "sirclo/groupproject/restapi/delivery/controller/auth"
	_productController "sirclo/groupproject/restapi/delivery/controller/product"
	_productCategoryController "sirclo/groupproject/restapi/delivery/controller/productCategory"
	_userController "sirclo/groupproject/restapi/delivery/controller/user"

	_authRepo "sirclo/groupproject/restapi/repository/auth"
	_productRepo "sirclo/groupproject/restapi/repository/product"
	_productCategoryRepo "sirclo/groupproject/restapi/repository/productCategory"
	_userRepo "sirclo/groupproject/restapi/repository/user"

	"github.com/labstack/echo/v4"
)

func main() {
	// initialize database connection
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := config.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize model
	authRepo := _authRepo.NewAuthRepository(db)
	userRepo := _userRepo.NewUserReposiroty(db)
	productRepo := _productRepo.NewProductRepositpry(db)
	productCategoryRepo := _productCategoryRepo.NewProductCategoryRepository(db)

	// initialize controller
	authController := _authController.NewAuthController(authRepo)
	userController := _userController.NewUserController(userRepo)
	productController := _productController.NewProductController(productRepo, userRepo)
	productCategoryController := _productCategoryController.NewProductCategoryController(productCategoryRepo)

	// create new echo
	e := echo.New()

	route.RegisterPath(e, authController, userController, productController, productCategoryController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
