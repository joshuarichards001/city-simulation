package protocol

type CommandType string

const (
	CommandTypeMove CommandType = "MOVE"
)

// Command represents a command to be sent to the browser.
type Command struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// MoveData represents the data for a citizen's movement.
type MoveCommandData struct {
	CitizenID int     `json:"citizenId"`
	FromX     float64 `json:"fromX"`
	FromY     float64 `json:"fromY"`
	ToX       float64 `json:"toX"`
	ToY       float64 `json:"toY"`
	Duration  int     `json:"duration"`
}