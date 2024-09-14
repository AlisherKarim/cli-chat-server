package handlers

import (
	"net/http"

	"github.com/alisherkarim/cli-chat-server/pkg/response"
)
	

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJson(w, 200, struct {}{})
}


func HandlerError(w http.ResponseWriter, r *http.Request) {
	response.RespondWithError(w, http.StatusOK, "Server Error")
}