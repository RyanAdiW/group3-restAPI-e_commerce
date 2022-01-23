package user

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
	"github.com/stretchr/testify/assert"
)

// 1. test create user
func TestCreateUser(t *testing.T) {
	t.Run("success create user", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "alta",
			"email":    "alta@mail.com",
			"password": "12345",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, userController.CreateUserController()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success create user", response.Message)
		}

	})
	t.Run("failed to bind data", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "alta",
			"email":    "alta@mail.com",
			"password": 12345,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, userController.CreateUserController()(context)) {
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
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "alta",
			"email":    "alta@mail.com",
			"password": "12345",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUserController(mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, userController.CreateUserController()(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "failed", response.Status)
			assert.Equal(t, "failed to create user", response.Message)
		}

	})
}

// 2. test get user by id
func TestGetUser(t *testing.T) {
	var (
		globalToken, errCreateToken = middlewares.CreateToken(2, "admin")
	)
	t.Run("success get user by id", func(t *testing.T) {
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := NewUserController(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUserController())(context)) {
			bodyResponses := res.Body.String()
			var response Responses

			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, 200, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "success get user", response.Message)
		}

	})
	t.Run("failed to convert id", func(t *testing.T) {
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("a")

		userController := NewUserController(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUserController())(context)) {
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
	t.Run("failed to getUserName from token", func(t *testing.T) {
		token, err := middlewares.CreateToken(2, "")
		if err != nil {
			panic(err)
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := NewUserController(mockUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUserController())(context)) {
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
	t.Run("error from repository", func(t *testing.T) {
		token, err := globalToken, errCreateToken
		if err != nil {
			panic(err)
		}
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := NewUserController(mockErrorUserRepository{})

		type Responses struct {
			Code    string `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUserController())(context)) {
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
