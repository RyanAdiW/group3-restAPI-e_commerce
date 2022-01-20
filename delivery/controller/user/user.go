package user

import (
	"net/http"
	response "sirclo/groupproject/restapi/delivery/common"
	"sirclo/groupproject/restapi/entities"
	userRepo "sirclo/groupproject/restapi/repository/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repository userRepo.UserRepository
}

func NewUserController(user userRepo.UserRepository) *UserController {
	return &UserController{repository: user}
}

// get user by id controller
func (uc UserController) GetUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get id from param
		userId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// get user from db
		user, err := uc.repository.GetUserById(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get user", user))
	}
}

// create new user	controller
func (uc UserController) CreateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// bind data
		var userRequest UserRequestFormat
		if err := c.Bind(&userRequest); err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}

		user := entities.User{
			Name:      userRequest.Name,
			User_name: userRequest.User_name,
			Email:     userRequest.Email,
			Password:  userRequest.Password,
			Born_date: userRequest.Born_date,
			Gender:    userRequest.Gender,
		}

		// create user to database
		err := uc.repository.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to create user"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create user"))
	}
}

// update user controller
func (uc UserController) UpdateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get id from param
		userId, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// binding data
		user := entities.User{}
		if errBind := c.Bind(&user); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}
		// update user based on id to database
		err := uc.repository.UpdateUser(user, userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "update success"))
	}
}

// delete user controller
func (uc UserController) DeleteUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get id from param
		userId, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// delete user based on id from database
		err := uc.repository.DeleteUser(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "delete success"))
	}
}
