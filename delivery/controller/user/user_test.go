package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sirclo/groupproject/restapi/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

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
