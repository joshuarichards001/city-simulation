package citizengenerator

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"

	"city-simulation/internal/citygenerator"
	"city-simulation/internal/config"
)

type BaseCitizen struct {
	ID    int
	Name  string
	HomeX int
	HomeY int
}

func GenerateCitizens(city *citygenerator.BaseCity) {
	buildingLocations := []struct{ x, y int }{}
	for i := range city.Grid {
		for j := range city.Grid[i] {
			if city.Grid[i][j] {
				buildingLocations = append(buildingLocations, struct{ x, y int }{i, j})
			}
		}
	}

	citizenNames := []string{
		"John Doe",
		"Jane Smith",
		"Alice Johnson",
		"Bob Brown",
		"Charlie Davis",
		"David Wilson",
		"Emily Taylor",
		"Frank Martin",
		"Grace Anderson",
		"Henry Thomas",
	}

	rand.Shuffle(len(buildingLocations), func(i, j int) {
		buildingLocations[i], buildingLocations[j] = buildingLocations[j], buildingLocations[i]
	})

	citizens := []BaseCitizen{}
	for i, name := range citizenNames {
		location := buildingLocations[i]
		citizens = append(citizens, BaseCitizen{
			ID:    i,
			Name:  name,
			HomeX: location.x,
			HomeY: location.y,
		})
	}

	file, err := os.Create(config.CitizensFilePath)
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
