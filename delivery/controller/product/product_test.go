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
func TestGetProduct(t *testing.T) {
	t.Run("success get all products", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductController(mockProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, productController.GetProductsController()(context)) {
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

		productController := NewProductController(mockErrorProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, productController.GetProductsController()(context)) {
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
			"stock": "qwewr",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductController(mockErrorProductRepository{}, mockErrorUserRepository{})

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
			assert.Equal(t, "failed to create data", response.Message)
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

		productController := NewProductController(mockErrorProductRepository{}, mockErrorUserRepository{})

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

// 3. test get by id
func TestGetById(t *testing.T) {
	t.Run("success get product by id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductController(mockProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, productController.GetByIdController()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, 200, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get product", response.Message)
		}

	})
	t.Run("failed to convert id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("a")

		productController := NewProductController(mockErrorProductRepository{}, mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, productController.GetByIdController()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to convert id", response.Message)
		}

	})
	t.Run("error from repository", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductController(mockErrorProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, productController.GetByIdController()(context)) {
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

// 4. test update product by id
func TestUpdateProduct(t *testing.T) {
	var (
		globalToken, errCreateToken = middlewares.CreateToken(1, "admin")
	)
	t.Run("failed convert id", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "buku",
			"price": 30000,
			"stock": 5,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("a")

		productController := NewProductController(mockErrorProductRepository{}, mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.UpdateProductController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to convert id", response.Message)
		}

	})
	t.Run("failed to get product by id", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "buku",
			"price": 30000,
			"stock": 5,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductController(mockErrorProductRepository{}, mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.UpdateProductController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to get product by id", response.Message)
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
			"price": 30000,
			"stock": 5,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductController(mockProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.UpdateProductController())(context)) {
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
	t.Run("user not belong to the product", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "buku",
			"price": 30000,
			"stock": 5,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductController(mockProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.UpdateProductController())(context)) {
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

// 5. test delete product
func TestDeleteProduct(t *testing.T) {
	var (
		globalToken, errCreateToken = middlewares.CreateToken(2, "admin")
	)
	t.Run("failed convert id", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("a")

		productController := NewProductController(mockErrorProductRepository{}, mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.DeleteProductController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to convert id", response.Message)
		}

	})
	t.Run("failed to get product by id", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductController(mockErrorProductRepository{}, mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.DeleteProductController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to get product by id", response.Message)
		}

	})
	t.Run("failed to get id from token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(-1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductController(mockProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.DeleteProductController())(context)) {
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
	t.Run("user not belong to the product", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductController(mockProductRepository{}, mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.DeleteProductController())(context)) {
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
	t.Run("error delete from repository", func(t *testing.T) {
		e := echo.New()
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")
		productController := NewProductController(mockOtherCaseProductRepository{}, mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middleware.JWT([]byte("rahasia"))(productController.DeleteProductController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "data not found", response.Message)
		}

	})
}

// mock product
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

func (m mockErrorProductRepository) GetProducts(idUser int) ([]entities.ProductResponseFormat, error) {
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

// mock user
type mockUserRepository struct{}

func (m mockUserRepository) CreateUser(entities.Users) error {
	return fmt.Errorf("error")
}
func (m mockUserRepository) GetUserById(id int) (entities.UserResponseFormat, error) {
	return entities.UserResponseFormat{}, nil
}
func (m mockUserRepository) UpdateUser(user entities.Users, id int) error {
	return fmt.Errorf("error")
}
func (m mockUserRepository) DeleteUser(id int) error {
	return fmt.Errorf("error")
}

type mockErrorUserRepository struct{}

func (m mockErrorUserRepository) CreateUser(entities.Users) error {
	return fmt.Errorf("error")
}
func (m mockErrorUserRepository) GetUserById(id int) (entities.UserResponseFormat, error) {
	return entities.UserResponseFormat{}, nil
}
func (m mockErrorUserRepository) UpdateUser(user entities.Users, id int) error {
	return fmt.Errorf("error")
}
func (m mockErrorUserRepository) DeleteUser(id int) error {
	return fmt.Errorf("error")
}

// mock other cases
type mockOtherCaseProductRepository struct{}

func (m mockOtherCaseProductRepository) GetProducts(idUser int) ([]entities.ProductResponseFormat, error) {
	return []entities.ProductResponseFormat{
		{Id: 1,
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
		{Id: 2,
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

func (m mockOtherCaseProductRepository) CreateProduct(entities.Products) error {
	return nil
}
func (m mockOtherCaseProductRepository) GetProductById(id int) (entities.ProductResponseFormat, error) {
	return entities.ProductResponseFormat{
		Id_user:             2,
		Username:            "tingkiwingki",
		Id_product_category: 1,
		Name_category:       "laptop",
		Name:                "MSI",
		Description:         "laptop kencang",
		Price:               20000000,
		Quantity:            5,
		Url_photo:           "msimodern.png",
	}, nil
}
func (m mockOtherCaseProductRepository) UpdateProduct(product entities.Products, id int) error {
	return fmt.Errorf("error")
}
func (m mockOtherCaseProductRepository) DeleteProduct(id int) error {
	return fmt.Errorf("error")
}
