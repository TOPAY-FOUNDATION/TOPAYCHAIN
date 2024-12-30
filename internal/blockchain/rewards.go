package blockchain

import (
	"math/big"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/wallet"
)

func (bc *Blockchain) DistributeRewards(miner string, reward *big.Int, tokenSymbol string) error {
	// Ensure the miner wallet exists
	minerWallet, exists := bc.Wallets[miner]
	if !exists {
		// Create a new wallet for the miner if it doesn't exist
		newWallet, err := wallet.NewWallet()
		if err != nil {
			return err
		}
		newWallet.Address = miner
		bc.Wallets[miner] = newWallet
		minerWallet = newWallet
	}

	// Ensure the wallet has a balance map for the token
	if _, exists := minerWallet.Balances[tokenSymbol]; !exists {
		minerWallet.Balances[tokenSymbol] = big.NewInt(0)
	}

	// Add the reward to the miner's balance
	minerWallet.Balances[tokenSymbol].Add(minerWallet.Balances[tokenSymbol], reward)

	return nil
}
