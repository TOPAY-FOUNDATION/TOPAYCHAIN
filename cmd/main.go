package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
	"tpy-blockchain/internal/blockchain"
	"tpy-blockchain/internal/common"
	"tpy-blockchain/internal/wallet"
)

func main() {
	// Initialize the blockchain
	bc := blockchain.NewBlockchain()
	fmt.Println("Blockchain initialized with genesis block.")

	// Create a UtilityToken using common.UtilityToken
	token := &common.UtilityToken{
		Name:        "TOPAY",
		Symbol:      "TPY",
		TotalSupply: big.NewInt(120000000),
		Decimals:    18,
		Balances:    make(map[string]*big.Int),
	}

	// Add the token to the blockchain
	err := bc.AddToken(token.Name, token.Symbol, token.TotalSupply, token.Decimals)
	if err != nil {
		fmt.Printf("Error adding token: %v\n", err)
	} else {
		fmt.Printf("Utility Token '%s' (%s) successfully added with total supply %s.\n", token.Name, token.Symbol, token.TotalSupply.String())
	}

	// Command-line interface loop
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- TOPAY Blockchain CLI ---")
		fmt.Println("1. Create Wallet")
		fmt.Println("2. View Blockchain")
		fmt.Println("3. Add Transaction")
		fmt.Println("4. View Wallet Balance")
		fmt.Println("5. Transfer Tokens")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		// Read user input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			handleCreateWallet(bc, token)
		case "2":
			handleViewBlockchain(bc)
		case "3":
			handleAddTransaction(bc, reader)
		case "4":
			handleViewWalletBalance(reader, token) // Removed bc
		case "5":
			handleTransferTokens(reader, token)    // Removed bc		
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

// Handle wallet creation
// Handle wallet creation
func handleCreateWallet(bc *blockchain.Blockchain, token *common.UtilityToken) {
	fmt.Println("\nCreating a new wallet...")
	w, err := wallet.NewWallet()
	if err != nil {
		fmt.Println("Error creating wallet:", err)
		return
	}

	// Add wallet to blockchain
	bc.Wallets[w.Address] = w

	// Assign initial token balance
	initialBalance := big.NewInt(1000)
	token.Balances[w.Address] = initialBalance

	// Create a new block to reflect the wallet creation
	block := &blockchain.Block{
		Index:        len(bc.Blocks),
		Timestamp:    time.Now().String(),
		Transactions: []*blockchain.Transaction{}, // No transactions yet
		Wallets:      bc.Wallets,                  // Updated wallets
		Tokens:       bc.Tokens,                   // Updated tokens
		Nonce:        0,
		PreviousHash: bc.Blocks[len(bc.Blocks)-1].Hash,
	}
	block.Hash = blockchain.CalculateHash(block)

	// Add the block to the blockchain
	bc.Blocks = append(bc.Blocks, block)

	// Save the updated blockchain
	err = bc.SaveBlocksToFile()
	if err != nil {
		fmt.Printf("Failed to save blockchain: %v\n", err)
		return
	}

	// Display wallet details
	fmt.Printf("Wallet Created Successfully!\nAddress: %s\nMnemonic: %s\n", w.Address, w.Mnemonic)
	fmt.Printf("Assigned initial balance of %s %s\n", initialBalance.String(), token.Symbol)
	fmt.Println("Blockchain updated with new block.")
}

// Handle viewing the blockchain
func handleViewBlockchain(bc *blockchain.Blockchain) {
	fmt.Println("\n--- Blockchain ---")
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("  Hash: %s\n", block.Hash)
		fmt.Printf("  Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("  Transactions: %d\n", len(block.Transactions))
	}
}

// Handle adding a transaction
func handleAddTransaction(bc *blockchain.Blockchain, reader *bufio.Reader) {
	fmt.Println("\nAdding a new transaction...")
	fmt.Print("Enter sender address: ")
	sender, _ := reader.ReadString('\n')
	sender = strings.TrimSpace(sender)

	fmt.Print("Enter receiver address: ")
	receiver, _ := reader.ReadString('\n')
	receiver = strings.TrimSpace(receiver)

	fmt.Print("Enter amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid amount, please try again.")
		return
	}

	transaction := &blockchain.Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   big.NewInt(amount),
	}
	err = bc.AddTransaction(transaction)
	if err != nil {
		fmt.Println("Error adding transaction:", err)
		return
	}

	fmt.Println("Transaction added successfully!")
}

// Handle viewing wallet balance
func handleViewWalletBalance(reader *bufio.Reader, token *common.UtilityToken) {
	fmt.Print("\nEnter wallet address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	balance := token.Balances[address]
	if balance == nil {
		balance = big.NewInt(0)
	}

	fmt.Printf("Wallet Balance: %s %s\n", balance.String(), token.Symbol)
}

// Handle transferring tokens
func handleTransferTokens(reader *bufio.Reader, token *common.UtilityToken) {
	fmt.Println("\nTransferring tokens...")
	fmt.Print("Enter sender address: ")
	sender, _ := reader.ReadString('\n')
	sender = strings.TrimSpace(sender)

	fmt.Print("Enter receiver address: ")
	receiver, _ := reader.ReadString('\n')
	receiver = strings.TrimSpace(receiver)

	fmt.Print("Enter amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid amount, please try again.")
		return
	}

	// Perform the token transfer
	err = token.Transfer(sender, receiver, big.NewInt(amount))
	if err != nil {
		fmt.Println("Error transferring tokens:", err)
		return
	}

	fmt.Println("Tokens transferred successfully!")
}
