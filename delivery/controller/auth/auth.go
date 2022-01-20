package auth

import (
	"net/http"
	"sirclo/groupproject/restapi/delivery/common"
	"sirclo/groupproject/restapi/repository/auth"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repository auth.Auth
}

func NewAuthController(repository auth.Auth) *AuthController {
	return &AuthController{repository: repository}
}

func (ac AuthController) LoginUserNameController() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginRequest LoginUserNameRequestFormat

		//bind request data
		if err := c.Bind(&loginRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("unauthorized", "failed to bind"))
		}

		// get token from login credential
		token, err := ac.repository.LoginUserName(loginRequest.User_name, loginRequest.Password)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("unauthorized", "user not found"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
			"name":  loginRequest.User_name,
		})
	}
}