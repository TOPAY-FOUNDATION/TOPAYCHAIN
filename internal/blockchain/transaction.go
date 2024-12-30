package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/wallet"
)

type Transaction struct {
	Sender      string
	Receiver    string
	Amount      *big.Int
	Signature   string
	Hash        string
	TokenSymbol string
}

func NewTransaction(senderWallet *wallet.Wallet, receiver string, amount *big.Int, tokenSymbol string) (*Transaction, error) {
	// Check sender's balance for the token
	balance, exists := senderWallet.Balances[tokenSymbol]
	if !exists || balance.Cmp(amount) < 0 {
		return nil, fmt.Errorf("insufficient balance for token %s", tokenSymbol)
	}

	// Create a transaction
	tx := &Transaction{
		Sender:      senderWallet.Address,
		Receiver:    receiver,
		Amount:      amount,
		TokenSymbol: tokenSymbol,
	}

	// Sign the transaction using the sender's wallet
	signature, err := senderWallet.Sign([]byte(tx.String()))
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}
	tx.Signature = signature

	return tx, nil
}

func (tx *Transaction) String() string {
	return fmt.Sprintf("%s:%s:%s:%s", tx.Sender, tx.Receiver, tx.Amount.String(), tx.TokenSymbol)
}

func (tx *Transaction) CalculateHash() string {
	record := fmt.Sprintf("%s:%s:%s:%s", tx.Sender, tx.Receiver, tx.Amount.String(), tx.TokenSymbol)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func (tx *Transaction) VerifySignature(publicKey *ecdsa.PublicKey) bool {
	hash := sha256.Sum256([]byte(tx.Hash))
	return ecdsa.Verify(publicKey, hash[:], new(big.Int).SetBytes([]byte(tx.Signature[:32])), new(big.Int).SetBytes([]byte(tx.Signature[32:])))
}
