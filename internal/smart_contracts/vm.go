package smart_contracts

import (
	"errors"
	"fmt"
)

type ExecutionContext struct {
	Sender    string
	Receiver  string
	Arguments map[string]interface{}
}

type VirtualMachine struct {
	GasLimit uint64
	GasUsed  uint64
}

func NewVirtualMachine(gasLimit uint64) *VirtualMachine {
	return &VirtualMachine{
		GasLimit: gasLimit,
		GasUsed:  0,
	}
}

func (vm *VirtualMachine) Execute(contract *SmartContract, function string, args map[string]interface{}) (interface{}, error) {
	vm.GasUsed += 10 // Example gas usage
	if vm.GasUsed > vm.GasLimit {
		return nil, errors.New("gas limit exceeded")
	}

	switch function {
	case "set":
		key, ok := args["key"].(string)
		if !ok {
			return nil, errors.New("missing or invalid key argument")
		}
		value := args["value"]
		contract.UpdateState(key, value)
		return "State updated", nil

	case "get":
		key, ok := args["key"].(string)
		if !ok {
			return nil, errors.New("missing or invalid key argument")
		}
		state, exists := contract.GetState(key)
		if !exists {
			return nil, fmt.Errorf("state for key '%s' not found", key)
		}
		return state, nil

	default:
		return nil, fmt.Errorf("function '%s' not recognized", function)
	}
}
