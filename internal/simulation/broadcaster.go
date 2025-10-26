package simulation

import (
	"city-simulation/internal/protocol"
	"encoding/json"
	"log"
)

// Broadcaster represents something that can send raw byte messages.
// The server's Hub type implements this interface.
type Broadcaster interface {
	Broadcast(message []byte)
}

// CommandBroadcaster adapts a byte-level Broadcaster to send structured protocol
// commands by JSON-encoding them first.
type CommandBroadcaster struct {
	out Broadcaster
}

func NewBroadcaster(out Broadcaster) *CommandBroadcaster {
	return &CommandBroadcaster{out: out}
}

func (b CommandBroadcaster) Broadcast(message []byte) {
	if b.out != nil {
		b.out.Broadcast(message)
	}
}

func (b *CommandBroadcaster) BroadcastCommand(cmd protocol.Command) {
	data, err := json.Marshal(cmd)
	if err != nil {
		log.Printf("Error marshaling command: %v", err)
		return
	}
	b.Broadcast(data)
}
