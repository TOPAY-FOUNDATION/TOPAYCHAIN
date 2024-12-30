package consensus

import (
	"crypto/rand"
	"math/big"
)

// GenerateRandomNumber generates a random number in the range [0, max).
func GenerateRandomNumber(max *big.Int) (*big.Int, error) {
	// Use crypto/rand to generate a secure random number
	randomNum, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, err
	}
	return randomNum, nil
}
