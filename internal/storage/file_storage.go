package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type FileStorage struct {
	BaseDir string
}

func NewFileStorage(baseDir string) *FileStorage {
	err := os.MkdirAll(baseDir, 0755)
	if err != nil {
		panic(fmt.Sprintf("Failed to create storage directory: %v", err))
	}
	return &FileStorage{BaseDir: baseDir}
}

func (fs *FileStorage) Save(key string, value interface{}) error {
	filePath := filepath.Join(fs.BaseDir, key+".json")
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}
	return os.WriteFile(filePath, data, 0644)
}

func (fs *FileStorage) Load(key string, result interface{}) error {
	filePath := filepath.Join(fs.BaseDir, key+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	return json.Unmarshal(data, result)
}

func (fs *FileStorage) Delete(key string) error {
	filePath := filepath.Join(fs.BaseDir, key+".json")
	return os.Remove(filePath)
}

func (fs *FileStorage) ListKeys() ([]string, error) {
	files, err := os.ReadDir(fs.BaseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}
	var keys []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			keys = append(keys, file.Name()[:len(file.Name())-5]) // Remove ".json" suffix
		}
	}
	return keys, nil
}
