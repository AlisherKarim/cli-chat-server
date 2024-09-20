package models

type User struct {
	Username string `json:"username"`
	Id int `json:"id,omitempty"`
	Email string `json:"email"`
	Password string `json:"-"`
}

type ChatRoom struct {
	Id string `json:"id"`
	Users []User `json:"users"`
}