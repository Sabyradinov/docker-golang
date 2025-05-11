package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelp_Get(t *testing.T) {
	// Set up Gin in test mode
	gin.SetMode(gin.TestMode)

	// Create a new Help handler
	helpHandler := NewHelpHandler()

	// Create a new Gin router and register the handler
	router := gin.Default()
	router.GET("/help", helpHandler.Get)

	// Create a test HTTP request
	req, _ := http.NewRequest(http.MethodGet, "/help", nil)

	// Create a test HTTP response recorder
	recorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(recorder, req)

	// Assert the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, recorder.Code)
	}

	// Assert the response body
	expectedBody := `"test"`
	if recorder.Body.String() != expectedBody {
		t.Errorf("expected body %s but got %s", expectedBody, recorder.Body.String())
	}
}
