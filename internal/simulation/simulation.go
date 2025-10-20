package simulation

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type Broadcaster interface {
	Broadcast(message []byte)
}

type Command struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type StartMoveData struct {
	CitizenID int     `json:"citizenId"`
	FromX     float64 `json:"fromX"`
	FromY     float64 `json:"fromY"`
	ToX       float64 `json:"toX"`
	ToY       float64 `json:"toY"`
	Duration  int     `json:"duration"`
}

type Citizen struct {
	ID        int     `json:"id"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	TargetX   float64 `json:"-"`
	TargetY   float64 `json:"-"`
	MoveUntil int64   `json:"-"`
}

type Simulation struct {
	broadcaster Broadcaster
	tickRate    time.Duration
	citizens    []Citizen
	stopSignal  chan bool
}

func NewSimulation(broadcaster Broadcaster) *Simulation {
	return &Simulation{
		broadcaster: broadcaster,
		tickRate:    100 * time.Millisecond,
		stopSignal:  make(chan bool),
		citizens:    generateCitizens(5),
	}
}

func generateCitizens(count int) []Citizen {
	citizens := make([]Citizen, count)
	for i := 0; i < count; i++ {
		x := rand.Float64() * 800
		y := rand.Float64() * 600
		citizens[i] = Citizen{
			ID:        i + 1,
			X:         x,
			Y:         y,
			TargetX:   x,
			TargetY:   y,
			MoveUntil: 0,
		}
	}
	return citizens
}

func (s *Simulation) Start() {
	log.Println("Simulation started")
	ticker := time.NewTicker(s.tickRate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.update()
		case <-s.stopSignal:
			log.Println("Simulation stopped")
			return
		}
	}
}

func (s *Simulation) Stop() {
	s.stopSignal <- true
}

func (s *Simulation) update() {
	now := time.Now().UnixMilli()

	for i := range s.citizens {
		if now >= s.citizens[i].MoveUntil {
			s.startNewMove(&s.citizens[i])
		}
	}
}

func (s *Simulation) startNewMove(citizen *Citizen) {
	citizen.X = citizen.TargetX
	citizen.Y = citizen.TargetY

	citizen.TargetX = rand.Float64() * 800
	citizen.TargetY = rand.Float64() * 600

	duration := rand.Intn(3000) + 2000

	citizen.MoveUntil = time.Now().UnixMilli() + int64(duration)

	cmd := Command{
		Type: "START_MOVE",
		Data: StartMoveData{
			CitizenID: citizen.ID,
			FromX:     citizen.X,
			FromY:     citizen.Y,
			ToX:       citizen.TargetX,
			ToY:       citizen.TargetY,
			Duration:  duration,
		},
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		log.Printf("Error marshaling command: %v", err)
		return
	}

	s.broadcaster.Broadcast(data)
}
