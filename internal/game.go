package internal

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type GameState struct {
	Tick       int       `json:"tick"`
	Population int       `json:"population"`
	Money      int       `json:"money"`
	Happiness  float64   `json:"happiness"`
	Citizens   []Citizen `json:"citizens"`
	Timestamp  int64     `json:"timestamp"`
}

type Citizen struct {
	ID   int     `json:"id"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Name string  `json:"name"`
}

type Game struct {
	hub        *Hub
	tickRate   time.Duration
	state      GameState
	stopSignal chan bool
}

func NewGame(hub *Hub) *Game {
	return &Game{
		hub:        hub,
		tickRate:   100 * time.Millisecond,
		stopSignal: make(chan bool),
		state: GameState{
			Tick:       0,
			Population: 5,
			Money:      10000,
			Happiness:  75.5,
			Citizens:   generateDummyCitizens(5),
		},
	}
}

func generateDummyCitizens(count int) []Citizen {
	citizens := make([]Citizen, count)
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}

	for i := 0; i < count; i++ {
		citizens[i] = Citizen{
			ID:   i + 1,
			X:    rand.Float64() * 800,
			Y:    rand.Float64() * 600,
			Name: names[i%len(names)],
		}
	}
	return citizens
}

func (g *Game) Start() {
	log.Println("Game loop started")
	ticker := time.NewTicker(g.tickRate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			g.update()
			g.broadcast()
		case <-g.stopSignal:
			log.Println("Game loop stopped")
			return
		}
	}
}

func (g *Game) Stop() {
	g.stopSignal <- true
}

func (g *Game) update() {
	g.state.Tick++
	g.state.Timestamp = time.Now().UnixMilli()

	for i := range g.state.Citizens {
		g.state.Citizens[i].X += (rand.Float64() - 0.5) * 5
		g.state.Citizens[i].Y += (rand.Float64() - 0.5) * 5

		g.state.Citizens[i].X = clamp(g.state.Citizens[i].X, 0, 800)
		g.state.Citizens[i].Y = clamp(g.state.Citizens[i].Y, 0, 600)
	}

	g.state.Money += rand.Intn(100) - 40
	g.state.Happiness += (rand.Float64() - 0.5) * 2
	g.state.Happiness = clamp(g.state.Happiness, 0, 100)
}

func (g *Game) broadcast() {
	data, err := json.Marshal(g.state)
	if err != nil {
		log.Printf("Error marshaling game state: %v", err)
		return
	}

	g.hub.Broadcast(data)
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
