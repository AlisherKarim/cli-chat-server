package controllers

import (
	"fmt"
	"strconv"

	"github.com/alisherkarim/cli-chat-server/db"
	"github.com/alisherkarim/cli-chat-server/ws"
)

type RoomController struct {
	rooms map[string]*ws.Hub
	storage db.Storage
}

func NewRoomController(storage db.Storage) *RoomController {
	return &RoomController{
		rooms: make(map[string]*ws.Hub),
		storage: storage,
	}
}

func (roomController *RoomController) AddRoom() (string, error) {
	newHub := ws.NewHub()
	hubId := strconv.Itoa(len(roomController.rooms) + 1)
	roomController.rooms[hubId] = newHub
	go newHub.Run()
	return hubId, nil
}

func (roomController *RoomController) GetRoom(id string) (*ws.Hub, error) {
	hub, found := roomController.rooms[id]
	if found {
		return hub, nil
	}
	return nil, fmt.Errorf("room with id %s not found", id)
}