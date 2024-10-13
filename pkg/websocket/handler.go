package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	log.Println("handle web socket connections")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
			log.Printf("Error upgrading connection: %v", err)
			return
	}
	defer conn.Close()

	data, err := json.Marshal("Hello from server")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Println(err)
		return
	}
}
