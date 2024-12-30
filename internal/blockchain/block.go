package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Index        int
	Timestamp    string
	Transactions []*Transaction
	Nonce        int
	PreviousHash string
	Hash         string
}

func CalculateHash(block *Block) string {
	record := fmt.Sprintf("%d%s%d%s", block.Index, block.Timestamp, block.Nonce, block.PreviousHash)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func (block *Block) MineBlock(difficulty int) {
	target := ""
	for i := 0; i < difficulty; i++ {
		target += "0"
	}

	for block.Hash[:difficulty] != target {
		block.Nonce++
		block.Hash = CalculateHash(block)
	}
	fmt.Printf("Block mined: %s\n", block.Hash)
}
