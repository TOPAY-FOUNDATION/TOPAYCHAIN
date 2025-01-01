package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/TOPAY-FOUNDATION/TOPAYCHAIN/internal/blockchain"
)

// APIServer defines the structure of the API server
type APIServer struct {
	Router      *mux.Router
	Port        string
	Blockchain  *blockchain.Blockchain
}

// NewAPIServer initializes a new API server
func NewAPIServer(bc *blockchain.Blockchain) *APIServer {
	port := os.Getenv("PORT") // Use PORT from environment variables
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	server := &APIServer{
		Router:     mux.NewRouter(),
		Port:       port,
		Blockchain: bc,
	}

	// Register routes with both the router and blockchain instance
	RegisterRoutes(server.Router, server.Blockchain)

	return server
}

// Start starts the API server
func (s *APIServer) Start() {
	log.Printf("Starting API server on port %s...", s.Port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+s.Port, s.Router)) // Listen on all interfaces
}
