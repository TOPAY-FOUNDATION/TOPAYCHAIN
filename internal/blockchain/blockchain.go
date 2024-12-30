package blockchain

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/common"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/smart_contracts"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/wallet"
)

type Blockchain struct {
	Blocks       []*Block
	Wallets      map[string]*wallet.Wallet
	Tokens       map[string]*common.UtilityToken
	Transactions []*Transaction
	mutex        sync.Mutex
}

func (bc *Blockchain) AddSmartContract(contract *smart_contracts.SmartContract) {
	panic("unimplemented")
}

func (bc *Blockchain) GetBalance(address string) (any, any) {
	panic("unimplemented")
}

func (bc *Blockchain) GetSmartContract(contractAddress string) (any, any) {
	panic("unimplemented")
}

func (bc *Blockchain) AddTransaction(tx *Transaction) any {
	panic("unimplemented")
}

func (bc *Blockchain) AddWallet(w *wallet.Wallet) {
	panic("unimplemented")
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
