package wallet

import "math/big"

type Asset struct {
	Type    string
	Symbol  string
	Balance *big.Int
}

func NewAsset(assetType, symbol string) *Asset {
	return &Asset{
		Type:    assetType,
		Symbol:  symbol,
		Balance: big.NewInt(0),
	}
}

func (a *Asset) UpdateBalance(amount *big.Int) {
	a.Balance.Add(a.Balance, amount)
}
