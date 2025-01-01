package api

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/blockchain"
)

// Blockchain instance
var blockchainInstance *blockchain.Blockchain

// RegisterRoutes registers all API endpoints
func RegisterRoutes(router *mux.Router, bc *blockchain.Blockchain) {
	blockchainInstance = bc

	router.HandleFunc("/blocks", handleGetBlocks).Methods("GET")
	router.HandleFunc("/blocks/mine", handleMineBlock).Methods("POST")
	router.HandleFunc("/wallets", handleCreateWallet).Methods("POST")
	router.HandleFunc("/wallets/{address}/balance", handleGetBalance).Methods("GET")
	router.HandleFunc("/transactions", handleAddTransaction).Methods("POST")
	router.HandleFunc("/transactions/pending", handleGetPendingTransactions).Methods("GET")
	router.HandleFunc("/rpc", handleJSONRPC).Methods("POST")
}

// handleGetBlocks retrieves all blocks
func handleGetBlocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if blockchainInstance == nil {
		http.Error(w, "Blockchain not initialized", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(blockchainInstance.GetBlocks())
}

// handleMineBlock mines a new block
func handleMineBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	block, err := blockchainInstance.MineBlock()
	if err != nil {
		if e, ok := err.(error); ok {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "Unknown error occurred", http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(block)
}

// handleCreateWallet creates a wallet
func handleCreateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	wallet, err := blockchainInstance.CreateWallet()
	if err != nil {
		if e, ok := err.(error); ok {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "Unknown error occurred", http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(wallet)
}

// handleGetBalance retrieves a wallet's balance
func handleGetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	address := vars["address"]
	balance, err := blockchainInstance.GetBalance(address)
	if err != nil {
		if e, ok := err.(error); ok {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "Unknown error occurred", http.StatusInternalServerError)
		}
		return
	}

	if bigBalance, ok := balance.(*big.Int); ok {
		json.NewEncoder(w).Encode(map[string]string{"address": address, "balance": bigBalance.String()})
	} else {
		http.Error(w, "Balance type mismatch", http.StatusInternalServerError)
	}
}

// handleAddTransaction creates a new transaction
func handleAddTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tx blockchain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := blockchainInstance.AddTransaction(&tx)
	if err != nil {
		if e, ok := err.(error); ok {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "Unknown error occurred", http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Transaction added successfully!"})
}

// handleGetPendingTransactions retrieves all pending transactions
func handleGetPendingTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pendingTransactions := blockchainInstance.GetPendingTransactions()
	json.NewEncoder(w).Encode(pendingTransactions)
}

// handleJSONRPC processes JSON-RPC requests
func handleJSONRPC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	method, _ := request["method"].(string)
	switch method {
	case "eth_blockNumber":
		handleBlockNumber(w)
	case "eth_getBalance":
		handleGetBalanceJSONRPC(w, request)
	default:
		http.Error(w, "Method not supported", http.StatusNotImplemented)
	}
}

// handleBlockNumber retrieves the current block number for JSON-RPC
func handleBlockNumber(w http.ResponseWriter) {
	blockNumber := blockchainInstance.GetCurrentBlockNumber()
	if blockNum, ok := blockNumber.(*big.Int); ok {
		json.NewEncoder(w).Encode(map[string]string{"result": blockNum.String()})
	} else {
		http.Error(w, "Block number type mismatch", http.StatusInternalServerError)
	}
}

// handleGetBalanceJSONRPC retrieves a wallet's balance for JSON-RPC
func handleGetBalanceJSONRPC(w http.ResponseWriter, request map[string]interface{}) {
	params, _ := request["params"].([]interface{})
	address, _ := params[0].(string)
	balance, err := blockchainInstance.GetBalance(address)
	if err != nil {
		if e, ok := err.(error); ok {
			http.Error(w, e.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "Unknown error occurred", http.StatusInternalServerError)
		}
		return
	}

	if bigBalance, ok := balance.(*big.Int); ok {
		json.NewEncoder(w).Encode(map[string]string{"result": bigBalance.String()})
	} else {
		http.Error(w, "Balance type mismatch", http.StatusInternalServerError)
	}
}
