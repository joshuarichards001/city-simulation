package citizengenerator

import (
	"encoding/json"
	"log"
	"os"
)

type BaseCitizen struct {
	ID      int
	Name    string
	HomeX   int
	HomeY   int
}

func GenerateCitizens() {
	citizens := []BaseCitizen{
		{ID: 1, Name: "John Doe", HomeX: 10, HomeY: 20},
		{ID: 2, Name: "Jane Smith", HomeX: 30, HomeY: 40},
		{ID: 3, Name: "Alice Johnson", HomeX: 50, HomeY: 60},
		{ID: 4, Name: "Bob Brown", HomeX: 70, HomeY: 80},
		{ID: 5, Name: "Charlie Davis", HomeX: 90, HomeY: 100},
		{ID: 6, Name: "David Wilson", HomeX: 110, HomeY: 120},
		{ID: 7, Name: "Emily Taylor", HomeX: 130, HomeY: 140},
		{ID: 8, Name: "Frank Martin", HomeX: 150, HomeY: 160},
		{ID: 9, Name: "Grace Anderson", HomeX: 170, HomeY: 180},
		{ID: 10, Name: "Henry Thomas", HomeX: 190, HomeY: 200},
	}

	file, err := os.Create("data/citizens.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(citizens)
	if err != nil {
		log.Fatal(err)
	}
}
