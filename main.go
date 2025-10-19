package main

import (
	"log"

	"city-simulation/internal"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	server := internal.NewServer()
	if err := server.Start(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
