package consensus

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

func CalculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func GenerateRandomNumber(max *big.Int) *big.Int {
	return big.NewInt(0).Rand(big.NewInt(0), max)
}

func VerifyBlockHash(hash string, difficulty int) bool {
	prefix := ""
	for i := 0; i < difficulty; i++ {
		prefix += "0"
	}
	return hash[:difficulty] == prefix
}
