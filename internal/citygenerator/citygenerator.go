package citygenerator

import (
	"encoding/json"
	"log"
	"os"
	
	"city-simulation/internal/config"
)

// BaseCity represents a city grid with a boolean grid, true values indicate buildings, false values indicate roads.
type BaseCity struct {
	Grid [][]bool `json:"grid"`
}

var CitySize int = 20

func GenerateCity() *BaseCity {
	city := &BaseCity{
		Grid: make([][]bool, CitySize),
	}

	for i := 0; i < CitySize; i++ {
		city.Grid[i] = make([]bool, CitySize)
		for j := 0; j < CitySize; j++ {
			if i%2 == 0 && j%2 == 0 {
				city.Grid[i][j] = true
			} else {
				city.Grid[i][j] = false
			}
		}
	}

	printCity(city)
	
	return city
}

func printCity(city *BaseCity) {
	file, err := os.Create(config.CityFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(city)
	if err != nil {
		log.Fatal(err)
	}
}
