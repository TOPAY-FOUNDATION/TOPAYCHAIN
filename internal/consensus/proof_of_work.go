package consensus

import (
	"fmt"
	"strings"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/blockchain"
)

type ProofOfWork struct {
	Difficulty int
}

func NewProofOfWork(difficulty int) *ProofOfWork {
	return &ProofOfWork{
		Difficulty: difficulty,
	}
}

func (pow *ProofOfWork) MineBlock(block *blockchain.Block) {
	target := strings.Repeat("0", pow.Difficulty)

	for block.Hash[:pow.Difficulty] != target {
		block.Nonce++
		block.Hash = blockchain.CalculateHash(block)
	}
	fmt.Printf("Block mined with hash: %s\n", block.Hash)
}

func (pow *ProofOfWork) ValidateBlock(block *blockchain.Block) bool {
	target := strings.Repeat("0", pow.Difficulty)
	return strings.HasPrefix(block.Hash, target)
}
