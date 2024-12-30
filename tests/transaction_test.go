package tests

import (
	"math/big"
	"testing"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/blockchain"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/wallet"
)

func TestNewTransaction(t *testing.T) {
	senderWallet, _ := wallet.NewWallet()
	receiver := "a123b456c789d012e345f678901234567890abcd"
	amount := big.NewInt(100)

	tx, err := blockchain.NewTransaction(senderWallet, receiver, amount, "TPY")
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	if tx.Receiver != receiver {
		t.Errorf("Transaction receiver mismatch. Expected %s, got %s", receiver, tx.Receiver)
	}
}
