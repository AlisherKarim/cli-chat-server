package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alisherkarim/cli-chat-server/pkg/response"
	"github.com/alisherkarim/cli-chat-server/ws"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


type RoomRequest struct {
	Name string `json:"name"`
}

func (mainHandler *MainHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var req RoomRequest

	// Parse JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("error decoding request body: %v", err)
		response.RespondWithErrorMsg(w, http.StatusBadRequest, "invalid request")
		return
	}

	// Use the parsed name
	id, err := mainHandler.roomController.AddRoom(req.Name)
	if err != nil {
		log.Printf("error: %v", err)
		response.RespondWithErrorMsg(w, 500, "server error")
		return
	}

	msg := fmt.Sprintf("created a room with id: %s", id)
	response.RespondWithJson(w, http.StatusCreated, msg)
}

func (mainHandler *MainHandler) HandleJoin(w http.ResponseWriter, r *http.Request) {
	log.Println("handle web socket connections")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
			log.Printf("Error upgrading connection: %v", err)
			return
	}
	defer conn.Close()
	
}

func (mainHandler *MainHandler) HandleGetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := mainHandler.roomController.GetRooms()
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// return in format {"rooms": [ {"id": "123", "name": "friends room"} ]}
	resp := map[string]interface{}{
		"rooms": rooms,
	}

	response.RespondWithJson(w, http.StatusOK, resp)
}

func (mainHandler *MainHandler) HandleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	log.Println("handle web socket connections")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
			log.Printf("Error upgrading connection: %v", err)
			return
	}
	defer conn.Close()
	
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	chatRoom, err := mainHandler.roomController.GetRoom(id)
	if err != nil {
		log.Printf("error: %v", err)
		conn.Close()
	}

	// create a client
	// start client listenings
	// add client to the hub

	client := ws.Client{
		Hub: chatRoom.Hub,
		Conn: conn,
		Send: make(chan []byte),
	}
	chatRoom.Hub.Register <- &client

	go client.WritePump()
	// client.Conn.WriteMessage(websocket.TextMessage, []byte("You were connected to room " + chatRoom.DataBaseModel.Name))
	client.ReadPump()
}
