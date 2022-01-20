package main

import (
	"os"
	"sirclo/groupproject/restapi/config"
	_userController "sirclo/groupproject/restapi/delivery/controller/user"
	"sirclo/groupproject/restapi/delivery/route"
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

	// initialize controller
	userController := _userController.NewUserController(userRepo)

	// create new echo
	e := echo.New()

	route.RegisterPath(e, userController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
