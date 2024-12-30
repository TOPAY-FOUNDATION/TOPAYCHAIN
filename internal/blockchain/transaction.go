package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Transaction struct {
	Sender      string
	Receiver    string
	Amount      *big.Int
	Signature   string
	Hash        string
	TokenSymbol string
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
