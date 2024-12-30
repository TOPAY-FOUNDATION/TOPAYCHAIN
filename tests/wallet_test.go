package tests

import (
	"testing"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/wallet"
)

func TestNewWallet(t *testing.T) {
	w, err := wallet.NewWallet()
	if err != nil {
		t.Fatalf("Failed to create wallet: %v", err)
	}

	if len(w.Address) != 40 {
		t.Errorf("Invalid wallet address length: %d", len(w.Address))
	}
}
