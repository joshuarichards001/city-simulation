package simulation

import (
	"log"
	"time"
)

// Simulation manages the simulation state and handles citizen movement.
type Simulation struct {
	broadcaster *CommandBroadcaster
	tickRate    time.Duration
	citizens    []*Citizen
	stopSignal  chan bool
}

func NewSimulation(broadcaster Broadcaster) *Simulation {
	simulation := &Simulation{
		broadcaster: NewBroadcaster(broadcaster),
		tickRate:    100 * time.Millisecond,
		stopSignal:  make(chan bool),
		citizens:    loadCitizens(),
	}

	return simulation
}

func (simulation *Simulation) Start() {
	log.Println("Simulation started")
	ticker := time.NewTicker(simulation.tickRate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			simulation.update()
		case <-simulation.stopSignal:
			log.Println("Simulation stopped")
			return
		}
	}
}

func (simulation *Simulation) Stop() {
	simulation.stopSignal <- true
}

func (simulation *Simulation) update() {
	now := time.Now().UnixMilli()

	for _, c := range simulation.citizens {
		if now >= c.MoveUntil {
			cmd := c.Move()
			simulation.broadcaster.BroadcastCommand(cmd)
		}
	}
}
