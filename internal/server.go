package internal

import (
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "8080"
	webDistPath = "./web/dist"
)

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
	game := NewGame(hub)

	s := &Server{
		port:   port,
		router: http.NewServeMux(),
	}

	staticHandler := NewStaticHandler(webDistPath)
	s.router.Handle("/", staticHandler)
	s.router.HandleFunc("/ws", HandleWebSocket(hub))

	go hub.Run()
	go game.Start()

	return s
}

func (s *Server) Start() error {
	addr := ":" + s.port
	log.Printf("Starting server on port %s, serving static files from %s", s.port, webDistPath)
	return http.ListenAndServe(addr, s.router)
}
