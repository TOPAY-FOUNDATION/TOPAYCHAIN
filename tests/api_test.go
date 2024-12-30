package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/api"
)

func TestGetBlocks(t *testing.T) {
	// Create a router and register routes
	router := mux.NewRouter()
	api.RegisterRoutes(router)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/blocks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request through the router
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Optionally, check the response body
	expected := `[{"index":0,"hash":"abc123","transactions":[]}]`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
	}
}
