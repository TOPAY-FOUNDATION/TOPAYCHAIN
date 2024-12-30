package blockchain

import "math/big"

func (bc *Blockchain) DistributeRewards(miner string, amount *big.Int) {
	if _, exists := bc.Wallets[miner]; !exists {
		bc.Wallets[miner] = &Wallet{
			Address: miner,
			Balance: big.NewInt(0),
		}
	}
	bc.Wallets[miner].Balance.Add(bc.Wallets[miner].Balance, amount)
}
