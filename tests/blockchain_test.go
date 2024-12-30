package tests

import (
	"testing"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/blockchain"
)

func TestAddBlock(t *testing.T) {
	bc := blockchain.NewBlockchain()
	previousBlockCount := len(bc.Blocks)

	err := bc.AddBlock([]*blockchain.Transaction{})
	if err != nil {
		t.Fatalf("Failed to add block: %v", err)
	}

	if len(bc.Blocks) != previousBlockCount+1 {
		t.Errorf("Block not added. Expected %d blocks, got %d", previousBlockCount+1, len(bc.Blocks))
	}
}
