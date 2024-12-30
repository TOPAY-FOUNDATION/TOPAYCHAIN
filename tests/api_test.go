package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/api"
)

func TestGetBlocks(t *testing.T) {
	req, err := http.NewRequest("GET", "/blocks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.HandleGetBlocks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}
