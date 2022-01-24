package order

import (
	"net/http"
	"strconv"

	response "sirclo/groupproject/restapi/delivery/common"
	middlewares "sirclo/groupproject/restapi/delivery/middleware"
	"sirclo/groupproject/restapi/entities"
	orderRepo "sirclo/groupproject/restapi/repository/order"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	repository orderRepo.Order
}

func NewOrderController(order orderRepo.Order) *OrderController {
	return &OrderController{repository: order}
}

func (oc OrderController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := middlewares.GetId("rahasia", c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get order from db
		orders, err := oc.repository.Get(userId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get user history order", orders))
	}
}

func (oc OrderController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// bind data
		var orderRequest OrderRequestFormat
		if err := c.Bind(&orderRequest); err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}
		userId, errGetId := middlewares.GetId("rahasia", c)

		if errGetId != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		address_delivery := entities.Address_delivery{
			Street: orderRequest.Address_delivery.Street,
			City:   orderRequest.Address_delivery.City,
			State:  orderRequest.Address_delivery.State,
			Zip:    orderRequest.Address_delivery.Zip,
		}

		credit_card := entities.Credit_card{
			Type:   orderRequest.Credit_card.Type,
			Name:   orderRequest.Credit_card.Name,
			Number: orderRequest.Credit_card.Number,
			Cvv:    orderRequest.Credit_card.Cvv,
			Month:  orderRequest.Credit_card.Month,
			Year:   orderRequest.Credit_card.Year,
		}

		order := entities.Order{
			Id_user:          userId,
			Status:           "PENDING",
			Total_price:      orderRequest.Total_price,
			Address_delivery: address_delivery,
			Credit_card:      credit_card,
			Id_cart:          orderRequest.Id_cart,
		}

		// create order to database
		err := oc.repository.Create(order)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to create order"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create order"))
	}
}

func (oc OrderController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := middlewares.GetUserName(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		// get id from param
		orderId, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// binding data
		order := entities.Order{}
		if errBind := c.Bind(&order); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}

		// update order based on id to database
		errUpdate := oc.repository.Update(order, orderId)
		if errUpdate != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update order status"))
	}
}
