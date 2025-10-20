package main

import (
	"log"

	"city-simulation/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	srv := server.NewServer()
	if err := srv.Start(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
