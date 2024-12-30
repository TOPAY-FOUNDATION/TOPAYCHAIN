package smart_contracts

import (
	"time"
)

type SmartContract struct {
	Address   string                 `json:"address"`
	Owner     string                 `json:"owner"`
	Code      []byte                 `json:"code"`      // Smart contract bytecode
	State     map[string]interface{} `json:"state"`     // Persistent contract state
	CreatedAt time.Time              `json:"createdAt"` // Contract creation timestamp
}

func NewSmartContract(address, owner string, code []byte) *SmartContract {
	return &SmartContract{
		Address:   address,
		Owner:     owner,
		Code:      code,
		State:     make(map[string]interface{}),
		CreatedAt: time.Now(),
	}
}

func (sc *SmartContract) UpdateState(key string, value interface{}) {
	sc.State[key] = value
}

func (sc *SmartContract) GetState(key string) (interface{}, bool) {
	value, exists := sc.State[key]
	return value, exists
}
