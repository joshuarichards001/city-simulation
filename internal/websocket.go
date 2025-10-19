package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
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
}

func NewWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocketHandler created")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Received message: %s", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
