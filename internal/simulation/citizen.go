package simulation

import (
	"math/rand"
	"time"
	
	"city-simulation/internal/protocol"
)

// Citizen represents a citizen in the simulation.
type Citizen struct {
	ID        int     `json:"id"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	TargetX   float64 `json:"-"`
	TargetY   float64 `json:"-"`
	MoveUntil int64   `json:"-"`
}

func NewCitizen(id int, x, y float64) *Citizen {
	return &Citizen{
		ID:        id,
		X:         x,
		Y:         y,
		TargetX:   x,
		TargetY:   y,
		MoveUntil: 0,
	}
}

func (citizen *Citizen) Move() protocol.Command {
	citizen.X = citizen.TargetX
	citizen.Y = citizen.TargetY

	citizen.TargetX = rand.Float64() * 800
	citizen.TargetY = rand.Float64() * 600

	duration := rand.Intn(3000) + 2000

	citizen.MoveUntil = time.Now().UnixMilli() + int64(duration)

	cmd := protocol.Command{
		Type: "MOVE",
		Data: protocol.MoveCommandData{
			CitizenID: citizen.ID,
			FromX:     citizen.X,
			FromY:     citizen.Y,
			ToX:       citizen.TargetX,
			ToY:       citizen.TargetY,
			Duration:  duration,
		},
	}

	return cmd
}
