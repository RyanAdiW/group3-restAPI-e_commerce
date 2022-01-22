package product

import (
	"fmt"
	"net/http"
	"strconv"

	response "sirclo/groupproject/restapi/delivery/common"
	middlewares "sirclo/groupproject/restapi/delivery/middleware"
	"sirclo/groupproject/restapi/entities"
	productRepo "sirclo/groupproject/restapi/repository/product"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	repository productRepo.ProductRepository
}

func NewProductController(product productRepo.ProductRepository) *ProductController {
	return &ProductController{repository: product}
}

// 1. get all products controller
func (pc ProductController) GetProductsController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// filter by id_user
		uid := c.QueryParam("uid")

		// default value for user id
		if uid == "" {
			uid = "0"
		}

		idUser, err := strconv.Atoi(uid)

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id_user"))
		}

		// get all products from database
		products, err := pc.repository.GetProducts(idUser)

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		if len(products) == 0 {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all products", products))
	}
}

// 2. get product by id controller
func (pc ProductController) GetByIdController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get id from parameter
		productId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// get product from database based on id
		product, err := pc.repository.GetProductById(productId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get product", product))
	}
}

// 3. create product controller
func (pc ProductController) CreateProductController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// request data binding
		var productRequest ProductRequestFormat
		if err := c.Bind(&productRequest); err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}
		// get user id from login token
		userId, errToken := middlewares.GetId("rahasia", c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}

		product := entities.Products{
			Id_product_category: productRequest.Id_product_category,
			Name:                productRequest.Name,
			Description:         productRequest.Description,
			Price:               productRequest.Price,
			Quantity:            productRequest.Quantity,
			Url_photo:           productRequest.Url_photo,
			Id_user:             userId,
		}
		// post product to database
		err := pc.repository.CreateProduct(product)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to create data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success create product"))
	}
}

// 4. update product controller
func (pc ProductController) UpdateProductController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get id from parameter
		productId, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// get product based on id to get user id
		product, err := pc.repository.GetProductById(productId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to get product by id"))
		}
		// get user id from login token
		userId, errToken := middlewares.GetId("rahasia", c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "Unauthorized access"))
		}
		// check if the product belong to the user
		if product.Id_user != userId {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}
		// request data binding for update
		productBind := entities.Products{}
		if errBind := c.Bind(&productBind); errBind != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to bind data"))
		}
		// update product to database
		errRepo := pc.repository.UpdateProduct(productBind, productId)
		if errRepo != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "success update data"))
	}
}

// 5. delete product
func (pc ProductController) DeleteProductController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get id from parameter
		productId, errConv := strconv.Atoi(c.Param("id"))
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to convert id"))
		}
		// get product based on id to get user id
		product, err := pc.repository.GetProductById(productId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to get product by id"))
		}
		// get user id from login token
		userId, errToken := middlewares.GetId("rahasia", c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}
		// check if the product belong to the user
		if product.Id_user != userId {
			return c.JSON(http.StatusUnauthorized, response.UnauthorizedRequest("unauthorized", "unauthorized access"))
		}
		// delete product from database based on id
		errRepo := pc.repository.DeleteProduct(productId)
		if errRepo != nil {
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "data not found"))
		}
		return c.JSON(http.StatusOK, response.SuccessOperationDefault("success", "delete success"))
	}
}
