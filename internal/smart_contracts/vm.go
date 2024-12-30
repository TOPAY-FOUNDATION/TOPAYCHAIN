package smart_contracts

import (
	"errors"
	"fmt"
)

type ExecutionContext struct {
	Sender    string                 // Address of the sender
	Receiver  string                 // Address of the contract
	Arguments map[string]interface{} // Arguments passed to the contract
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
	// Decode and simulate contract execution
	vm.GasUsed += 10 // Example gas usage per execution
	if vm.GasUsed > vm.GasLimit {
		return nil, errors.New("gas limit exceeded")
	}

	// Example: Simulate contract logic with JSON interpretation
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
		return contract.GetState(key)

	default:
		return nil, fmt.Errorf("function %s not recognized", function)
	}
}
