package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all API endpoints with the given router
func RegisterRoutes(router *mux.Router) {
	// Blockchain routes
	router.HandleFunc("/blocks", handleGetBlocks).Methods("GET")
	router.HandleFunc("/blocks/mine", handleMineBlock).Methods("POST")

	// Wallet routes
	router.HandleFunc("/wallets", handleCreateWallet).Methods("POST")
	router.HandleFunc("/wallets/{address}/balance", handleGetBalance).Methods("GET")

	// Transaction routes
	router.HandleFunc("/transactions", handleAddTransaction).Methods("POST")
	router.HandleFunc("/transactions/pending", handleGetPendingTransactions).Methods("GET")
}

// handleGetBlocks handles the API request to retrieve all blocks
func handleGetBlocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Example response (Replace with actual blockchain data retrieval logic)
	blocks := []map[string]interface{}{
		{"index": 0, "hash": "abc123", "transactions": []string{}},
	}
	json.NewEncoder(w).Encode(blocks)
}

// handleMineBlock handles the API request to mine a new block
func handleMineBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Example response (Replace with actual block mining logic)
	json.NewEncoder(w).Encode(map[string]string{"message": "Block mined successfully!"})
}

// handleCreateWallet handles the API request to create a new wallet
func handleCreateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Example response (Replace with actual wallet creation logic)
	json.NewEncoder(w).Encode(map[string]string{"address": "tpy1qxyz123", "mnemonic": "example mnemonic"})
}

// handleGetBalance handles the API request to retrieve a wallet's balance
func handleGetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	address := vars["address"]
	// Example response (Replace with actual balance retrieval logic)
	json.NewEncoder(w).Encode(map[string]string{"address": address, "balance": "1000"})
}

// handleAddTransaction handles the API request to add a new transaction
func handleAddTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Example response (Replace with actual transaction creation logic)
	json.NewEncoder(w).Encode(map[string]string{"message": "Transaction added successfully!"})
}

// handleGetPendingTransactions handles the API request to retrieve pending transactions
func handleGetPendingTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Example response (Replace with actual pending transactions retrieval logic)
	pendingTransactions := []map[string]interface{}{
		{"sender": "tpy1qabc", "receiver": "tpy1qxyz", "amount": "100"},
	}
	json.NewEncoder(w).Encode(pendingTransactions)
}
