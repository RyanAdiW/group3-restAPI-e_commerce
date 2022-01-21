package user

import (
	"net/http"
	"strconv"

	"sirclo/groupproject/restapi/entities"

	response "sirclo/groupproject/restapi/delivery/common"
	middlewares "sirclo/groupproject/restapi/delivery/middleware"
	userRepo "sirclo/groupproject/restapi/repository/user"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repository userRepo.UserRepository
}

func NewUserController(user userRepo.UserRepository) *UserController {
	return &UserController{repository: user}
}

// 1. get user by id controller
func (uc UserController) GetUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUserName, err := middlewares.GetUserName(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

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

		return c.JSON(http.StatusOK, response.SuccessOperationLogin("success", "success get user", user, currentUserName))
	}
}

// 2. create new user	controller
func (uc UserController) CreateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// bind data
		var userRequest UserRequestFormat
		if err := c.Bind(&userRequest); err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}

		user := entities.Users{
			Name:       userRequest.Name,
			Username:   userRequest.Username,
			Email:      userRequest.Email,
			Password:   userRequest.Password,
			Birth_date: userRequest.Birth_date,
			Gender:     userRequest.Gender,
			Url_photo:  userRequest.Url_photo,
		}

		// create user to database
		err := uc.repository.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to create user"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create user"))
	}
}

// 3. update user controller
func (uc UserController) UpdateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := middlewares.GetUserName(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		userId, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// binding data
		user := entities.Users{}
		if errBind := c.Bind(&user); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}
		// update user based on id to database
		errUpdate := uc.repository.UpdateUser(user, userId)
		if errUpdate != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update user"))
	}
}

// 4. delete user controller
func (uc UserController) DeleteUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := middlewares.GetUserName(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		userId, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// delete user based on id from database
		errDelete := uc.repository.DeleteUser(userId)
		if errDelete != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "delete success"))
	}
}
