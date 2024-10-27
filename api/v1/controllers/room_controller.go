package controllers

import (
	"fmt"

	"github.com/alisherkarim/cli-chat-server/db"
	types "github.com/alisherkarim/cli-chat-server/internal"
	"github.com/alisherkarim/cli-chat-server/models"
	"github.com/alisherkarim/cli-chat-server/ws"
)

type RoomController struct {
	rooms map[string]types.ChatRoom
	storage db.Storage
}

func NewRoomController(storage db.Storage) *RoomController {
	return &RoomController{
		rooms: make(map[string]types.ChatRoom),
		storage: storage,
	}
}

func (roomController *RoomController) AddRoom(name string) (string, error) {
	newHub := ws.NewHub()
	room, err := roomController.storage.CreateRoom(name)
	if err != nil {
		return "", err
	}
	roomController.rooms[room.Id] = types.ChatRoom{
		Hub: newHub,
		DataBaseModel: room,
	}
	go newHub.Run()
	return room.Id, nil
}

func (roomController *RoomController) GetRoom(id string) (types.ChatRoom, error) {
	chatRoom, found := roomController.rooms[id]

	if found {
		return chatRoom, nil
	}

	dbRoom, err := roomController.storage.GetRoom(id)
	if err != nil {
		return types.ChatRoom{}, fmt.Errorf("failed to get room with id: %s", id)
	}

	newHub := ws.NewHub()
	go newHub.Run()

	newMapItem := types.ChatRoom{
		Hub: newHub,
		DataBaseModel: dbRoom,
	}
	roomController.rooms[dbRoom.Id] = newMapItem
	return roomController.rooms[dbRoom.Id], nil
}

func (roomController *RoomController) GetRooms() ([]models.ChatRoom, error) {
	rooms, err := roomController.storage.GetRooms()

	if err != nil {
		return []models.ChatRoom{}, err
	}
	return rooms, nil
}