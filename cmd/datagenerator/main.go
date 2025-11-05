package main

import (
	"city-simulation/internal/citygenerator"
	"city-simulation/internal/citizengenerator"
)

func main() {
	city := citygenerator.GenerateCity()
	citizengenerator.GenerateCitizens(city)
}
