package wallet

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

func RecoverWallet(mnemonic string) (*Wallet, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("invalid mnemonic")
	}

	seed := bip39.NewSeed(mnemonic, "")
	privateKey, err := crypto.ToECDSA(seed[:32])
	if err != nil {
		return nil, err
	}

	publicKey := &privateKey.PublicKey
	address := crypto.PubkeyToAddress(*publicKey).Hex()

	return &Wallet{
		Mnemonic:   mnemonic,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address,
		Balances:   make(map[string]*big.Int),
	}, nil
}
