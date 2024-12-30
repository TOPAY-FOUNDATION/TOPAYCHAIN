package tests

import (
	"math/big"
	"testing"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/consensus"
)

func TestProofOfStake(t *testing.T) {
	pos := consensus.NewProofOfStake()
	pos.AddStake("validator1", big.NewInt(1000))

	validator, err := pos.SelectValidator()
	if err != nil {
		t.Fatalf("Error selecting validator: %v", err)
	}

	if validator != "validator1" {
		t.Errorf("Expected validator1 to be selected, got %s", validator)
	}
}
