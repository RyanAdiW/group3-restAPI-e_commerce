package auth

import (
	"net/http"
	"sirclo/groupproject/restapi/delivery/common"
	"sirclo/groupproject/restapi/repository/auth"

	"golang.org/x/crypto/bcrypt"

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
		password := []byte(loginRequest.Password)

		hashedPassword, errPass := ac.repository.GetPasswordByUsername(loginRequest.Username)
		if errPass != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("unauthorized", "user not found"))
		}

		errMatch := bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
		if errMatch != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("unauthorized", "user not found"))
		}

		// get token from login credential
		token, err := ac.repository.LoginUserName(loginRequest.Username, hashedPassword)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("unauthorized", "user not found"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
			"name":  loginRequest.Username,
		})
	}
}
