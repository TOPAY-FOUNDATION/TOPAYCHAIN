package storage

import (
	"fmt"
	"sync"
)

type Cache struct {
	store map[string]interface{}
	mutex sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]interface{}),
	}
}

func (c *Cache) Save(key string, value interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.store[key] = value
	return nil
}

func (c *Cache) Load(key string, result interface{}) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, exists := c.store[key]
	if !exists {
		return fmt.Errorf("key not found: %s", key)
	}
	*result.(*interface{}) = value
	return nil
}

func (c *Cache) Delete(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.store, key)
	return nil
}

func (c *Cache) ListKeys() ([]string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	keys := make([]string, 0, len(c.store))
	for key := range c.store {
		keys = append(keys, key)
	}
	return keys, nil
}
