package handlers

import (
	"github.com/alisherkarim/cli-chat-server/api/v1/controllers"
	"github.com/alisherkarim/cli-chat-server/db"
)

type MainHandler struct {
	userController controllers.UserController
	roomController controllers.RoomController
}

func NewHandler(storage db.Storage) *MainHandler {
	return &MainHandler{
		userController: *controllers.NewUserController(storage),
		roomController: *controllers.NewRoomController(storage),
	}
}
