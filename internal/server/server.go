package server

import (
	"log"
	"net/http"
	"os"

	"city-simulation/internal/simulation"
)

const (
	defaultPort = "8080"
	webDistPath = "./web/dist"
)

// Server serves static files, manages routing, and handles WebSocket connections.
type Server struct {
	port   string
	router *http.ServeMux
}

func NewServer() *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	hub := NewHub()
	simulation := simulation.NewSimulation(hub)
	server := &Server{
		port:   port,
		router: http.NewServeMux(),
	}
	staticHandler := NewStaticHandler(webDistPath)
	websocketHandler := NewWebSocketHandler(hub)

	server.router.Handle("/", staticHandler)
	server.router.Handle("/ws", websocketHandler)

	go hub.Run()
	go simulation.Start()

	return server
}

func (server *Server) Start() error {
	addr := ":" + server.port
	log.Printf("Starting server on port %s, serving static files from %s", server.port, webDistPath)
	return http.ListenAndServe(addr, server.router)
}
