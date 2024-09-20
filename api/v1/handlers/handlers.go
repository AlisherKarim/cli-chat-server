package handlers

import (
	"net/http"

	"github.com/alisherkarim/cli-chat-server/api/v1/controllers"
	"github.com/alisherkarim/cli-chat-server/db"
	"github.com/alisherkarim/cli-chat-server/pkg/response"
)

type MainHandler struct {
	userController controllers.UserController
}

func NewHandler(storage db.Storage) *MainHandler {
	return &MainHandler{userController: *controllers.NewUserController(storage)}
}

func (mainHandler *MainHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	user, err := mainHandler.userController.GetUserById(r.URL.Query().Get("id"))

	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	response.RespondWithJson(w, http.StatusOK, user)
}

func (mainHandler *MainHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := mainHandler.userController.GetUsers()
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	response.RespondWithJson(w, http.StatusOK, users)
}