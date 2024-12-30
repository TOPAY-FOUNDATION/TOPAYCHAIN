package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/blockchain"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/smart_contracts"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/wallet"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/pkg/utils"
)

func main() {
	logger := utils.NewLogger()
	logger.Info("Starting TOPAYCHAIN Blockchain...")

	// Initialize blockchain
	bc := blockchain.NewBlockchain()
	fmt.Println("Blockchain initialized with genesis block.")

	// CLI loop
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- TOPAYCHAIN CLI ---")
		fmt.Println("1. Create Wallet")
		fmt.Println("2. View Blockchain")
		fmt.Println("3. Add Transaction")
		fmt.Println("4. View Wallet Balance")
		fmt.Println("5. Deploy Smart Contract")
		fmt.Println("6. Execute Smart Contract")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handleCreateWallet(bc)
		case "2":
			handleViewBlockchain(bc)
		case "3":
			handleAddTransaction(bc, reader)
		case "4":
			handleViewWalletBalance(bc, reader)
		case "5":
			handleDeploySmartContract(bc, reader)
		case "6":
			handleExecuteSmartContract(bc, reader)
		case "7":
			fmt.Println("Exiting TOPAYCHAIN CLI. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func handleCreateWallet(bc *blockchain.Blockchain) {
	w, err := wallet.NewWallet()
	if err != nil {
		fmt.Printf("Error creating wallet: %v\n", err)
		return
	}

	bc.AddWallet(w)
	fmt.Printf("Wallet Created Successfully!\nAddress: %s\nMnemonic: %s\n", w.Address, w.Mnemonic)
}

func handleViewBlockchain(bc *blockchain.Blockchain) {
	fmt.Println("\n--- Blockchain ---")
	for _, block := range bc.Blocks {
		fmt.Printf("Block %d | Hash: %s | Previous Hash: %s | Transactions: %d\n", block.Index, block.Hash, block.PreviousHash, len(block.Transactions))
	}
}

func handleAddTransaction(bc *blockchain.Blockchain, reader *bufio.Reader) {
	fmt.Print("Enter sender address: ")
	sender, _ := reader.ReadString('\n')
	sender = strings.TrimSpace(sender)

	fmt.Print("Enter receiver address: ")
	receiver, _ := reader.ReadString('\n')
	receiver = strings.TrimSpace(receiver)

	fmt.Print("Enter amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount := big.NewInt(0)
	amount.SetString(amountStr, 10)

	tx := &blockchain.Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}

	err := bc.AddTransaction(tx)
	if err != nil {
		fmt.Printf("Failed to add transaction: %v\n", err)
		return
	}

	fmt.Println("Transaction added successfully!")
}

func handleViewWalletBalance(bc *blockchain.Blockchain, reader *bufio.Reader) {
	fmt.Print("Enter wallet address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	balance, err := bc.GetBalance(address)
	if err != nil {
		fmt.Printf("Error retrieving balance: %v\n", err)
		return
	}

	// Ensure `balance` is of type `*big.Int`
	if bigBalance, ok := balance.(*big.Int); ok {
		fmt.Printf("Wallet Balance: %s\n", bigBalance.String())
	} else {
		fmt.Println("Error: Balance is not of expected type *big.Int")
	}
}

func handleDeploySmartContract(bc *blockchain.Blockchain, reader *bufio.Reader) {
	fmt.Print("Enter contract owner address: ")
	owner, _ := reader.ReadString('\n')
	owner = strings.TrimSpace(owner)

	address := smart_contracts.GenerateContractAddress()
	contract := smart_contracts.NewSmartContract(address, owner, []byte("example_code"))

	bc.AddSmartContract(contract)
	fmt.Printf("Smart Contract Deployed Successfully!\nAddress: %s\nOwner: %s\n", address, owner)
}

func handleExecuteSmartContract(bc *blockchain.Blockchain, reader *bufio.Reader) {
	fmt.Print("Enter contract address: ")
	contractAddress, _ := reader.ReadString('\n')
	contractAddress = strings.TrimSpace(contractAddress)

	contract, err := bc.GetSmartContract(contractAddress)
	if err != nil {
		fmt.Printf("Error fetching contract: %v\n", err)
		return
	}

	// Ensure `contract` is of type `*smart_contracts.SmartContract`
	sc, ok := contract.(*smart_contracts.SmartContract)
	if !ok {
		fmt.Println("Error: Contract is not of expected type *smart_contracts.SmartContract")
		return
	}

	fmt.Print("Enter function name (e.g., set, get): ")
	function, _ := reader.ReadString('\n')
	function = strings.TrimSpace(function)

	args := make(map[string]interface{})
	fmt.Print("Enter arguments as key=value pairs (comma-separated): ")
	argStr, _ := reader.ReadString('\n')
	argStr = strings.TrimSpace(argStr)

	for _, pair := range strings.Split(argStr, ",") {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			args[kv[0]] = kv[1]
		}
	}

	vm := smart_contracts.NewVirtualMachine(100000)
	result, err := vm.Execute(sc, function, args) // Use `sc` instead of `contract`
	if err != nil {
		fmt.Printf("Error executing contract: %v\n", err)
		return
	}

	fmt.Printf("Smart Contract Execution Result: %v\n", result)
}
