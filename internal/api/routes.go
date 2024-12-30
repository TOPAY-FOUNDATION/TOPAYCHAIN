package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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

func handleGetBlocks(w http.ResponseWriter, r *http.Request) {
	// Example response
	blocks := []map[string]interface{}{
		{"index": 0, "hash": "abc123", "transactions": []string{}},
	}
	json.NewEncoder(w).Encode(blocks)
}

func handleMineBlock(w http.ResponseWriter, r *http.Request) {
	// Example response
	json.NewEncoder(w).Encode(map[string]string{"message": "Block mined successfully!"})
}

func handleCreateWallet(w http.ResponseWriter, r *http.Request) {
	// Example response
	json.NewEncoder(w).Encode(map[string]string{"address": "tpy1qxyz123", "mnemonic": "example mnemonic"})
}

func handleGetBalance(w http.ResponseWriter, r *http.Request) {
	// Example response
	vars := mux.Vars(r)
	address := vars["address"]
	json.NewEncoder(w).Encode(map[string]string{"address": address, "balance": "1000"})
}

func handleAddTransaction(w http.ResponseWriter, r *http.Request) {
	// Example response
	json.NewEncoder(w).Encode(map[string]string{"message": "Transaction added successfully!"})
}

func handleGetPendingTransactions(w http.ResponseWriter, r *http.Request) {
	// Example response
	pendingTransactions := []map[string]interface{}{
		{"sender": "tpy1qabc", "receiver": "tpy1qxyz", "amount": "100"},
	}
	json.NewEncoder(w).Encode(pendingTransactions)
}
