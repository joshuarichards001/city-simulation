package internal

import (
	"log"
	"net/http"
	"os"
)

const (
	defaultPort      = "8080"
	frontendDistPath = "./frontend/dist"
)

type Server struct {
	port   string
	router *http.ServeMux
}

func New() *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	s := &Server{
		port:   port,
		router: http.NewServeMux(),
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	staticHandler := NewStaticHandler(frontendDistPath)
	s.router.Handle("/", staticHandler)
}

func (s *Server) Start() error {
	addr := ":" + s.port
	log.Printf("Starting server on port %s, serving static files from %s", s.port, frontendDistPath)
	return http.ListenAndServe(addr, s.router)
}
