package tests

import (
	"testing"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/storage"
)

func TestFileStorage(t *testing.T) {
	fs := storage.NewFileStorage("./testdata")

	err := fs.Save("testKey", map[string]string{"key": "value"})
	if err != nil {
		t.Fatalf("Failed to save data: %v", err)
	}

	var result map[string]string
	err = fs.Load("testKey", &result)
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if result["key"] != "value" {
		t.Errorf("Data mismatch. Expected 'value', got '%s'", result["key"])
	}
}
