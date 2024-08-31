package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/muditsaxena1/url-shortner/internal/api"
	"github.com/muditsaxena1/url-shortner/internal/models"
	"github.com/muditsaxena1/url-shortner/internal/storage"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	api.SetupRoutes(router)
	return router
}

func TestSetUser(t *testing.T) {
	router := setupRouter()

	// user := models.User{ID: "1", Name: "John Doe", SignupTime: 1622548800000} // Example Unix millisecond timestamp

	tests := []struct {
		name       string
		user       models.User
		statusCode int
		response   string
	}{
		{
			name:       "missing id field",
			user:       models.User{Name: "John Wick", SignupTime: 1622548800000},
			statusCode: http.StatusBadRequest,
			response:   `{"errors":{"id":"id is required"}}`,
		},
		{
			name:       "missing name field",
			user:       models.User{ID: "123", SignupTime: 1622548800000},
			statusCode: http.StatusBadRequest,
			response:   `{"errors":{"name":"Name is required"}}`,
		},
		{
			name:       "name too short",
			user:       models.User{ID: "123", Name: "J", SignupTime: 1622548800000},
			statusCode: http.StatusBadRequest,
			response:   `{"errors":{"name":"name must be at least 2 characters long"}}`,
		},
		{
			name:       "missing signupTime field",
			user:       models.User{ID: "123", Name: "John Wick"},
			statusCode: http.StatusBadRequest,
			response:   `{"errors":{"signupTime":"Signup time is required"}}`,
		},
		{
			name:       "signupTime before 1850",
			user:       models.User{ID: "123", Name: "John Wick", SignupTime: -3786825700000},
			statusCode: http.StatusBadRequest,
			response:   `{"errors":{"signupTime":"signupTime must be after 1 Jan 1850"}}`,
		},
		{
			name:       "valid input",
			user:       models.User{ID: "123", Name: "John Wick", SignupTime: 1622548800000},
			statusCode: http.StatusOK,
			response:   `{"status":"user set successfully"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonValue, _ := json.Marshal(tt.user)
			req, _ := http.NewRequest("POST", "/v1/user", bytes.NewBuffer(jsonValue))
			req.SetBasicAuth("admin", "password")
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			assert.Equal(t, tt.statusCode, resp.Code)
			assert.JSONEq(t, tt.response, resp.Body.String())
		})
	}

}

func TestGetUser(t *testing.T) {
	router := setupRouter()

	// Seed the storage with a user via API
	seedUser := models.User{ID: "123", Name: "John Wick", SignupTime: 1622548800000}
	api.UsersStorage = storage.NewInMemoryUserStorage()
	api.UsersStorage.SetUser(seedUser)

	tests := []struct {
		name       string
		userID     string
		statusCode int
		response   string
	}{
		{
			name:       "user exists",
			userID:     "123",
			statusCode: http.StatusOK,
			response:   `{"id":"123","name":"John Wick","signupTime":1622548800000}`,
		},
		{
			name:       "user does not exist",
			userID:     "999",
			statusCode: http.StatusNotFound,
			response:   `{"error":"user not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/v1/user/"+tt.userID, nil)
			req.SetBasicAuth("admin", "password")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.statusCode, w.Code)
			assert.JSONEq(t, tt.response, w.Body.String())
		})
	}
}

func TestListUsers(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name       string
		statusCode int
		response   string
		setupFunc  func()
	}{
		{
			name:       "multiple users exist",
			statusCode: http.StatusOK,
			response: `[
				{"id":"123","name":"John Wick","signupTime":1622548800000},
				{"id":"456","name":"Jane Doe","signupTime":1622635200000}
			]`,
			setupFunc: func() {
				// Seed the storage with users via API
				users := []models.User{
					{ID: "123", Name: "John Wick", SignupTime: 1622548800000},
					{ID: "456", Name: "Jane Doe", SignupTime: 1622635200000},
				}
				api.UsersStorage = storage.NewInMemoryUserStorage()
				for _, user := range users {
					api.UsersStorage.SetUser(user)
				}
			},
		},
		{
			name:       "no users exist",
			statusCode: http.StatusOK,
			response:   `[]`,
			setupFunc: func() {
				api.UsersStorage = storage.NewInMemoryUserStorage()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run the setup function to prepare the test case
			tt.setupFunc()

			req, _ := http.NewRequest("GET", "/v1/users", nil)
			req.SetBasicAuth("admin", "password")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.statusCode, w.Code)
			assert.JSONEq(t, tt.response, w.Body.String())
		})
	}
}
