package consensus

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Validator struct {
	Address string
	Stake   *big.Int
}

type ProofOfStake struct {
	Validators map[string]*Validator
}

func NewProofOfStake() *ProofOfStake {
	return &ProofOfStake{
		Validators: make(map[string]*Validator),
	}
}

func (pos *ProofOfStake) AddStake(address string, amount *big.Int) {
	if _, exists := pos.Validators[address]; !exists {
		pos.Validators[address] = &Validator{
			Address: address,
			Stake:   big.NewInt(0),
		}
	}
	pos.Validators[address].Stake.Add(pos.Validators[address].Stake, amount)
}

func (pos *ProofOfStake) SelectValidator() (string, error) {
	totalStake := big.NewInt(0)
	for _, validator := range pos.Validators {
		totalStake.Add(totalStake, validator.Stake)
	}

	if totalStake.Cmp(big.NewInt(0)) == 0 {
		return "", fmt.Errorf("no stake available to select a validator")
	}

	// Use crypto/rand to generate a random number in the range [0, totalStake)
	target, err := rand.Int(rand.Reader, totalStake)
	if err != nil {
		return "", fmt.Errorf("failed to generate random number: %v", err)
	}

	accumulated := big.NewInt(0)
	for _, validator := range pos.Validators {
		accumulated.Add(accumulated, validator.Stake)
		if accumulated.Cmp(target) >= 0 {
			return validator.Address, nil
		}
	}

	return "", fmt.Errorf("validator selection failed")
}
