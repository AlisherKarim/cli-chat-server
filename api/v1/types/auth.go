package types

import "github.com/alisherkarim/cli-chat-server/models"

type RegisterRequestBody struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password" validate:"required,min=8,max=64,passwordComplexity"` // lets just add validate for now
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseBody struct {
	models.User
	AccessToken string `json:"access_token"`
}