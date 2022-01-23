package cart

import (
	"net/http"

	response "sirclo/groupproject/restapi/delivery/common"
	middlewares "sirclo/groupproject/restapi/delivery/middleware"
	"sirclo/groupproject/restapi/entities"
	cartRepo "sirclo/groupproject/restapi/repository/cart"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	repository cartRepo.Cart
}

func NewCartController(cart cartRepo.Cart) *CartController {
	return &CartController{repository: cart}
}

func (cc CartController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := middlewares.GetId("rahasia", c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get user from db
		carts, err := cc.repository.Get(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get user cart", carts))
	}
}

func (cc CartController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// bind data
		var cartRequest CartRequestFormat
		if err := c.Bind(&cartRequest); err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}
		userId, errGetId := middlewares.GetId("rahasia", c)

		if errGetId != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		cart := entities.Cart{
			Id_user:    userId,
			Id_product: cartRequest.Id_product,
			Quantity:   cartRequest.Quantity,
		}

		// create user to database
		err := cc.repository.Create(cart)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to create user cart"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create user cart"))
	}
}