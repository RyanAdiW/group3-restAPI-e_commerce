package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	middlewares "sirclo/groupproject/restapi/delivery/middleware"
	"sirclo/groupproject/restapi/entities"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

// 1. test get product

// 2. test create product
func TestCreateProduct(t *testing.T) {
	var (
		globalToken, errCreateToken = middlewares.CreateToken(2, "admin")
	)
	t.Run("success create product", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "buku",
			"price": 20000,
			"stock": 5,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductController(mockProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.CreateProductController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success create product", response.Message)
		}

	})
	t.Run("failed to bind data", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "buku",
			"price": "20000",
			"stock": 5,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductController(mockProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.CreateProductController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to bind data", response.Message)
		}

	})
	t.Run("error from repository", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "buku",
			"price": 20000,
			"stock": 5,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductController(mockProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.CreateProductController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to create product", response.Message)
		}

	})
	t.Run("failed to get id from token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "buku",
			"price": 20000,
			"stock": 5,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductController(mockProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.CreateProductController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusUnauthorized, res.Code)
			assert.Equal(t, "unauthorized", response.Status)
			assert.Equal(t, "unauthorized access", response.Message)
		}

	})

}

type mockProductRepository struct{}

func (m mockProductRepository) CreateProduct(entities.Products) error {
	return nil
}

func (m mockProductRepository) GetProducts(idUser int) ([]entities.ProductResponseFormat, error) {
	return []entities.ProductResponseFormat{
		{
			Id:                  1,
			Id_user:             1,
			Username:            "dipssy",
			Id_product_category: 1,
			Name_category:       "laptop",
			Name:                "macbook",
			Description:         "laptop keren",
			Price:               20000000,
			Quantity:            5,
			Url_photo:           "mackbookpro.png",
		},
		{
			Id:                  2,
			Id_user:             2,
			Username:            "tingkiwingki",
			Id_product_category: 1,
			Name_category:       "laptop",
			Name:                "MSI",
			Description:         "laptop kencang",
			Price:               20000000,
			Quantity:            5,
			Url_photo:           "msimodern.png",
		},
	}, nil
}

func (m mockProductRepository) GetProductById(id int) (entities.ProductResponseFormat, error) {
	return entities.ProductResponseFormat{}, nil
}
func (m mockProductRepository) UpdateProduct(user entities.Products, id int) error {
	return nil
}
func (m mockProductRepository) DeleteProduct(id int) error {
	return nil
}

type mockErrorProductRepository struct{}

func (m mockProductRepository) GetProduct() ([]entities.ProductResponseFormat, error) {
	return []entities.ProductResponseFormat{
		{
			Id:                  1,
			Id_user:             1,
			Username:            "dipssy",
			Id_product_category: 1,
			Name_category:       "laptop",
			Name:                "macbook",
			Description:         "laptop keren",
			Price:               20000000,
			Quantity:            5,
			Url_photo:           "mackbookpro.png",
		},
		{
			Id:                  2,
			Id_user:             2,
			Username:            "tingkiwingki",
			Id_product_category: 1,
			Name_category:       "laptop",
			Name:                "MSI",
			Description:         "laptop kencang",
			Price:               20000000,
			Quantity:            5,
			Url_photo:           "msimodern.png",
		},
	}, fmt.Errorf("error")
}

func (m mockErrorProductRepository) CreateProduct(entities.Products) error {
	return fmt.Errorf("error")
}
func (m mockErrorProductRepository) GetProductById(id int) (entities.ProductResponseFormat, error) {
	return entities.ProductResponseFormat{}, fmt.Errorf("error")
}
func (m mockErrorProductRepository) UpdateProduct(product entities.Products, id int) error {
	return fmt.Errorf("error")
}
func (m mockErrorProductRepository) DeleteProduct(id int) error {
	return fmt.Errorf("error")
}

type mockUserRepository struct{}

func (m mockUserRepository) CreateUser(entities.Users) error {
	return nil
}
func (m mockUserRepository) GetUserById(id int) (entities.UserResponseFormat, error) {
	return entities.UserResponseFormat{}, nil
}
func (m mockUserRepository) UpdateUser(user entities.Users, id int) error {
	return nil
}
func (m mockUserRepository) DeleteUser(id int) error {
	return nil
}

type mockErrorUserRepository struct{}

func (m mockErrorUserRepository) CreateUser(entities.Users) error {
	return fmt.Errorf("error")
}
func (m mockErrorUserRepository) GetUserById(id int) (entities.UserResponseFormat, error) {
	return entities.UserResponseFormat{}, fmt.Errorf("error")
}
func (m mockErrorUserRepository) UpdateUser(user entities.Users, id int) error {
	return fmt.Errorf("error")
}
func (m mockErrorUserRepository) DeleteUser(id int) error {
	return fmt.Errorf("error")
}
