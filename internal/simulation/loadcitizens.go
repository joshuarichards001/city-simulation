package simulation

import (
	"encoding/json"
	"log"
	"os"
	
	"city-simulation/internal/config"
)

func loadCitizens() []*Citizen {
	loadedCitizens, err := loadCitizensFromFile(config.CitizensFilePath)
	if err != nil {
		log.Printf("Failed to load citizens from file: %v; generating random citizens instead", err)
		return []*Citizen{}
	}
	
	return loadedCitizens
}

func loadCitizensFromFile(path string) ([]*Citizen, error) {
	type fileCitizen struct {
		ID    int
		Name  string
		HomeX int
		HomeY int
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var fileCitizens []fileCitizen
	if err := json.Unmarshal(content, &fileCitizens); err != nil {
		return nil, err
	}

	citizens := make([]*Citizen, 0, len(fileCitizens))
	for _, c := range fileCitizens {
		x := float64(c.HomeX)
		y := float64(c.HomeY)
		newcitizen := NewCitizen(c.ID, x, y)
		citizens = append(citizens, newcitizen)
	}

	return citizens, nil
}