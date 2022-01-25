package productCategory

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"sirclo/groupproject/restapi/entities"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetGetProductCategory(t *testing.T) {
	t.Run("success get all products", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductCategoryController(mockProductCategoryRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, productController.GetProductCategoryController()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, 200, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get all products", response.Message)
		}

	})
	t.Run("failed to get product from repository", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductCategoryController(mockErrorProductCategoryRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, productController.GetProductCategoryController()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to fetch data", response.Message)
		}
	})
}

type mockProductCategoryRepository struct{}

func (m mockProductCategoryRepository) GetProductCategory() ([]entities.ProductCategory, error) {
	return []entities.ProductCategory{
		{Id: 1,
			Name_category: "makanan",
		},
		{Id: 2,
			Name_category: "minuman",
		},
	}, nil
}

type mockErrorProductCategoryRepository struct{}

func (m mockErrorProductCategoryRepository) GetProductCategory() ([]entities.ProductCategory, error) {
	return []entities.ProductCategory{
		{Id: 1,
			Name_category: "makanan",
		},
		{Id: 2,
			Name_category: "minuman",
		},
	}, fmt.Errorf("error")
}
