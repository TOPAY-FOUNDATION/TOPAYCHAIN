package consensus

import (
	"math/big"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/blockchain"
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

func (pos *ProofOfStake) SelectValidator() string {
	totalStake := big.NewInt(0)
	for _, validator := range pos.Validators {
		totalStake.Add(totalStake, validator.Stake)
	}

	target := big.NewInt(0).Rand(big.NewInt(0), totalStake)
	accumulated := big.NewInt(0)

	for _, validator := range pos.Validators {
		accumulated.Add(accumulated, validator.Stake)
		if accumulated.Cmp(target) >= 0 {
			return validator.Address
		}
	}
	return ""
}

func (pos *ProofOfStake) ValidateBlock(block *blockchain.Block, validatorAddress string) bool {
	if _, exists := pos.Validators[validatorAddress]; !exists {
		return false
	}
	return blockchain.CalculateHash(block) == block.Hash
}
