package ws

import "log"

// Hub maintains the set of active clients and broadcasts messages to them.
type Hub struct {
    Clients    map[*Client]bool  // Registered clients
    Broadcast  chan []byte       // Inbound messages from the clients
    Register   chan *Client      // Register requests from the clients
    Unregister chan *Client      // Unregister requests from clients
}

// NewHub creates a new Hub.
func NewHub() *Hub {
    return &Hub{
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan []byte),
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
    }
}

// Run starts the hub and listens for incoming events (register, unregister, and broadcast).
func (h *Hub) Run() {
    for {
        select {
        case client := <-h.Register:
            h.Clients[client] = true
        case client := <-h.Unregister:
            if _, ok := h.Clients[client]; ok {
                delete(h.Clients, client)
                close(client.Send)
            }
        case message := <-h.Broadcast:
            processedMessage, err := ProcessMessage(message)

            if err != nil {
                log.Printf("error happened whil processing message: %v", err)
                return
            }

            for client := range h.Clients {
                select {
                case client.Send <- []byte(processedMessage.Content):
                default:
                    close(client.Send)
                    delete(h.Clients, client)
                }
            }
        }
    }
}
