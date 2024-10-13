package websocket

import (
	"encoding/json"
	"log"
)

// Message represents the structure of a WebSocket message.
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
	Type    string `json:"type"`  // e.g., "text", "image", "notification"
}

// ProcessMessage processes a raw message, applies validation or custom logic.
func ProcessMessage(rawMessage []byte) (*Message, error) {
	var message Message
	err := json.Unmarshal(rawMessage, &message)
	if err != nil {
		log.Printf("Error parsing message: %v", err)
		return nil, err
	}

	// Apply custom logic here, like filtering content or routing based on type
	return &message, nil
}
