package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alisherkarim/cli-chat-server/api/v1/types"
	"github.com/alisherkarim/cli-chat-server/pkg/response"
	"golang.org/x/crypto/bcrypt"
)



func (mainHandler *MainHandler) Login(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJson(w, http.StatusOK, "Login success")
}

func (mainHandler *MainHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	req := &types.RegisterRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	// hash the password here
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := mainHandler.userController.CreateUser(req.Username, req.Email, string(bytes))
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	response.RespondWithJson(w, http.StatusOK, user)
}