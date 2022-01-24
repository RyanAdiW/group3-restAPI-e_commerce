package productCategory

import (
	"fmt"
	"net/http"

	"sirclo/groupproject/restapi/repository/productCategory"

	response "sirclo/groupproject/restapi/delivery/common"

	"github.com/labstack/echo/v4"
)

type ProductCategoryController struct {
	repository productCategory.ProductCategoryRepository
}

func NewProductCategoryController(repository productCategory.ProductCategoryRepository) *ProductCategoryController {
	return &ProductCategoryController{repository: repository}
}

// get all product category controller
func (pcr ProductCategoryController) GetProductCategoryController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get all products from database
		products, err := pcr.repository.GetProductCategory()

		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, response.BadRequest("failed", "failed to fetch data"))
		}

		return c.JSON(http.StatusOK, response.SuccessOperation("success", "success get all products", products))
	}
}
