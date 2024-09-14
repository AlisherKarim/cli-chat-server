package handlers

import (
	"net/http"

	"github.com/alisherkarim/cli-chat-server/pkg/response"
)


func GetUsers(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Username string `json:"username"`
		Id string `json:"id"`
	}

	type data struct {
		Users []User `json:"users"`
	}

	response.RespondWithJson(w, http.StatusOK, data {
		Users: []User {
			{
				Username: "alisher",
				Id: "8k52sk2do",
			}, 
			{
				Username: "another user",
				Id: "123ksdhfsp23",
			},
		},
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Handler for creating a user
}
