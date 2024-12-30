package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port        string `json:"port"`
	DatabaseURI string `json:"databaseUri"`
	LogLevel    string `json:"logLevel"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}
	return &config, nil
}
