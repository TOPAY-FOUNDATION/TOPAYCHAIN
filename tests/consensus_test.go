package tests

import (
	"testing"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/consensus"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/blockchain"
)

func TestProofOfStake(t *testing.T) {
	pos := consensus.NewProofOfStake()
	pos.AddStake("validator1", big.NewInt(1000))

	validator := pos.SelectValidator()
	if validator != "validator1" {
		t.Errorf("Expected validator1 to be selected, got %s", validator)
	}
}
