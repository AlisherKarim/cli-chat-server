package models

type User struct {
	Username string `json:"username"`
	Id string `json:"id"`
	Email string `json:"email"`
}

type ChatRoom struct {
	Id string `json:"id"`
	Users []User `json:"users"`
}