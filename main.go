package main

import (
	"os"
	"sirclo/groupproject/restapi/config"
	"sirclo/groupproject/restapi/delivery/route"

	_authController "sirclo/groupproject/restapi/delivery/controller/auth"
	_userController "sirclo/groupproject/restapi/delivery/controller/user"

	_authRepo "sirclo/groupproject/restapi/repository/auth"
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
	userRepo := _userRepo.NewUserReposiroty(db)
	authRepo := _authRepo.NewAuthRepository(db)

	// initialize controller
	userController := _userController.NewUserController(userRepo)
	authController := _authController.NewAuthController(authRepo)

	// create new echo
	e := echo.New()

	route.RegisterPath(e, authController, userController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
