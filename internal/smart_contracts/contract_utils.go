package smart_contracts

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
)

func GenerateContractAddress() string {
	id := uuid.New().String()
	hash := sha256.Sum256([]byte(id))
	return hex.EncodeToString(hash[:])[:40] // Return the first 40 characters as the address
}

func ValidateArguments(args map[string]interface{}, requiredKeys []string) error {
	for _, key := range requiredKeys {
		if _, exists := args[key]; !exists {
			return fmt.Errorf("missing required argument: %s", key)
		}
	}
	return nil
}
