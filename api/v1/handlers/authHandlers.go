package handlers

import (
	"log"
	"net/http"

	"github.com/alisherkarim/cli-chat-server/pkg/response"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Print("Sign up request")
	response.RespondWithJson(w, http.StatusOK, "Login success")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	log.Print("Sign up request")
	response.RespondWithJson(w, http.StatusOK, "User signed up successfully")
}