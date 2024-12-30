package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/tyler-smith/go-bip39"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	Mnemonic   string
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    string
	Balances   map[string]*big.Int // Token balances
}

func NewWallet() (*Wallet, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	publicKey := &privateKey.PublicKey
	address := crypto.PubkeyToAddress(*publicKey).Hex()

	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, fmt.Errorf("failed to generate entropy: %v", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, fmt.Errorf("failed to generate mnemonic: %v", err)
	}

	return &Wallet{
		Mnemonic:   mnemonic,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address,
		Balances:   make(map[string]*big.Int),
	}, nil
}

func (w *Wallet) TransferTokens(receiver string, tokenSymbol string, amount *big.Int) error {
	balance, exists := w.Balances[tokenSymbol]
	if !exists || balance.Cmp(amount) < 0 {
		return errors.New("insufficient funds")
	}

	w.Balances[tokenSymbol].Sub(balance, amount)
	// Logic to update the receiver's balance on the blockchain
	return nil
}

func (w *Wallet) Sign(data []byte) (string, error) {
	hash := sha256.Sum256(data)
	signature, err := crypto.Sign(hash[:], w.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}
	return hex.EncodeToString(signature), nil
}
