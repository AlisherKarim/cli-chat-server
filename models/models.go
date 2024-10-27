package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `json:"username"`
	Id string `json:"id,omitempty"`
	Email string `json:"email"`
	Password string `json:"-"`
}

type Token struct {
	Username string
	Email string
	UserId string
	*jwt.StandardClaims
}

type ChatRoom struct {
	Id string `json:"room_id"`
	Name string `json:"name"`
}