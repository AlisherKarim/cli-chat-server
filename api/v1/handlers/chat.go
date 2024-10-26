package handlers

import (
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
}

func (mainHandler *MainHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	id, err := mainHandler.roomController.AddRoom()
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
	hub, err := mainHandler.roomController.GetRoom(id)
	if err != nil {
		log.Printf("error: %v", err)
		conn.Close()
	}

	// create a client
	// start client listenings
	// add client to the hub

	client := ws.Client{
		Hub: hub,
		Conn: conn,
		Send: make(chan []byte),
	}
	hub.Register <- &client

	go client.WritePump()
	client.ReadPump()
}
