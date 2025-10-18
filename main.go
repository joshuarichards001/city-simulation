package main

import (
	"log"

	"city-simulation/internal"
)

func main() {
	srv := internal.New()
	if err := srv.Start(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}