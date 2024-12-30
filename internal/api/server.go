package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Router *mux.Router
	Port   string
}

func NewAPIServer(port string) *APIServer {
	server := &APIServer{
		Router: mux.NewRouter(),
		Port:   port,
	}

	server.initializeRoutes()
	return server
}

func (s *APIServer) Start() {
	log.Printf("Starting server on port %s...", s.Port)
	log.Fatal(http.ListenAndServe(":"+s.Port, s.Router))
}

func (s *APIServer) initializeRoutes() {
	// Use routes from the routes.go file
	RegisterRoutes(s.Router)
}
