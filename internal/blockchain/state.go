package blockchain

import (
	"math/big"
)

type State struct {
	Balances map[string]*big.Int
}

func NewState() *State {
	return &State{
		Balances: make(map[string]*big.Int),
	}
}

func (s *State) GetBalance(address string) *big.Int {
	if balance, exists := s.Balances[address]; exists {
		return balance
	}
	return big.NewInt(0)
}

func (s *State) UpdateBalance(address string, amount *big.Int) error {
	if _, exists := s.Balances[address]; !exists {
		s.Balances[address] = big.NewInt(0)
	}
	s.Balances[address].Add(s.Balances[address], amount)
	return nil
}
