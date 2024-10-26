package ws

import (
	"github.com/gorilla/websocket"
)

// Client represents a single WebSocket connection.
type Client struct {
    Hub  *Hub             // Reference to the hub (manages all clients)
    Conn *websocket.Conn  // WebSocket connection
    Send chan []byte      // Channel to send messages to the client
}

// ReadPump listens for incoming messages from the WebSocket connection.
func (c *Client) ReadPump() {
    defer func() {
        c.Hub.Unregister <- c  // Unregister the client when the connection closes
        c.Conn.Close()
    }()

    for {
        _, message, err := c.Conn.ReadMessage()
        if err != nil {
            break
        }
        c.Hub.Broadcast <- message  // Broadcast the message to all other clients via the hub
    }
}

// WritePump listens for messages to send to the WebSocket connection.
func (c *Client) WritePump() {
    defer func() {
        c.Conn.Close()
    }()

    for {
        select {
        case message, ok := <-c.Send:
            if !ok {
                // The hub closed the channel.
                c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }

            c.Conn.WriteMessage(websocket.TextMessage, message)
        }
    }
}
