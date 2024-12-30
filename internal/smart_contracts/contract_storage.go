package smart_contracts

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ContractStorage struct {
	StorageDir string
}

func NewContractStorage(dir string) *ContractStorage {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(fmt.Sprintf("Failed to create storage directory: %v", err))
	}
	return &ContractStorage{StorageDir: dir}
}

func (cs *ContractStorage) SaveContract(contract *SmartContract) error {
	filePath := filepath.Join(cs.StorageDir, contract.Address+".json")
	data, err := json.MarshalIndent(contract, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize contract: %v", err)
	}
	return os.WriteFile(filePath, data, 0644)
}

func (cs *ContractStorage) LoadContract(address string) (*SmartContract, error) {
	filePath := filepath.Join(cs.StorageDir, address+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read contract file: %v", err)
	}
	var contract SmartContract
	if err := json.Unmarshal(data, &contract); err != nil {
		return nil, fmt.Errorf("failed to deserialize contract: %v", err)
	}
	return &contract, nil
}
