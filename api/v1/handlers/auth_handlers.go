package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alisherkarim/cli-chat-server/pkg/response"
)

func (mainHandler *MainHandler) Login(w http.ResponseWriter, r *http.Request) {
	log.Print("Sign up request")
	response.RespondWithJson(w, http.StatusOK, "Login success")
}

func (mainHandler *MainHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	req := &CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response.RespondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := mainHandler.userController.CreateUser(req.Username, req.Email, "")
	if err != nil {
		errorMessage, err := json.Marshal(struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		})

		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}

		response.RespondWithError(w, http.StatusInternalServerError, string(errorMessage))
		return
	}
	response.RespondWithJson(w, http.StatusOK, user)
}