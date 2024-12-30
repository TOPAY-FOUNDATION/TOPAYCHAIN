package blockchain

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"
	"tpy-blockchain/internal/common"
	"tpy-blockchain/internal/wallet"
)

type Blockchain struct {
	Blocks       []*Block
	Wallets      map[string]*wallet.Wallet
	Tokens       map[string]*common.UtilityToken
	Transactions []*Transaction
	mutex        sync.Mutex
}

func NewBlockchain() *Blockchain {
	genesisBlock := &Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []*Transaction{},
		Nonce:        0,
		PreviousHash: "0",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)

	return &Blockchain{
		Blocks:  []*Block{genesisBlock},
		Wallets: make(map[string]*wallet.Wallet),
		Tokens:  make(map[string]*common.UtilityToken),
	}
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) error {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{
		Index:        len(bc.Blocks),
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		Nonce:        0,
		PreviousHash: previousBlock.Hash,
	}
	newBlock.Hash = CalculateHash(newBlock)
	bc.Blocks = append(bc.Blocks, newBlock)
	return nil
}

func (bc *Blockchain) SaveToFile(filePath string) error {
	data, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
