package handlers

import (
	"fmt"
	"net/http"

	"github.com/alisherkarim/cli-chat-server/pkg/response"
)
	

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJson(w, http.StatusOK, struct {}{})
}

type CustomError struct {
	StatusCode int
	Err error
}

func (r *CustomError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func HandlerError(w http.ResponseWriter, r *http.Request) {
	response.RespondWithError(w, http.StatusOK, &CustomError{StatusCode: 500})
}