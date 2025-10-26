package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	hub      *Hub
	upgrader websocket.Upgrader
}

func NewWebSocketHandler(hub *Hub) *WebSocketHandler {
	return &WebSocketHandler{
		hub: hub,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")
				environment := os.Getenv("ENVIRONMENT")
				if environment == "development" {
					return true
				} else {
					return origin == "https://city-simulation.fly.dev"
				}
			},
		},
	}
}

func (handler *WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := handler.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := NewClient(handler.hub, conn)
	handler.hub.register <- client

	go client.writePump()
	go client.readPump()
}
