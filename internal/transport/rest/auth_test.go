package httphandler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/ZaiPeeKann/auth-service_pg/internal/service"
	mocks "github.com/ZaiPeeKann/auth-service_pg/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSingUp(t *testing.T) {
	testTable := []struct {
		Name               string
		InputBody          string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			Name:               "Ok",
			InputBody:          `{"username":"Something123", "email":"mymail@puregrade.com", "password":"mysecret"}`,
			expectedStatusCode: 200,
			expectedResponse:   `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		auth := mocks.NewAuthService()
		services := &service.Service{Authorization: auth}
		handler := NewHTTPHandler(services)
		r := gin.New() // router
		r.POST("/sing-up", handler.singUp)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/sing-up", bytes.NewBufferString(testCase.InputBody))

		r.ServeHTTP(w, req)

		assert.Equal(t, testCase.expectedStatusCode, w.Code)
		assert.Equal(t, testCase.expectedResponse, w.Body.String())
	}
}

func TestSingIn(t *testing.T) {
	testTable := []struct {
		Name               string
		InputBody          string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			Name:               "Ok",
			InputBody:          `{"username":"Something123", "password":"mysecret"}`,
			expectedStatusCode: 200,
			expectedResponse:   `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM2MDAwMDB9.IRGJFwYepzr7sK8trcdGUV9UF2Fw3FUHOXP-ktyHsZs"}`,
		},
	}

	for _, testCase := range testTable {
		auth := mocks.NewAuthService()
		services := &service.Service{Authorization: auth}
		handler := NewHTTPHandler(services)
		r := gin.New() // router
		r.POST("/sing-in", handler.singIn)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/sing-in", bytes.NewBufferString(testCase.InputBody))

		r.ServeHTTP(w, req)

		assert.Equal(t, testCase.expectedStatusCode, w.Code)
		assert.Equal(t, testCase.expectedResponse, w.Body.String())
	}
}
