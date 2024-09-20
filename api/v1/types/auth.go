package types

type RegisterRequestBody struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password" validate:"required,min=8,max=64,passwordComplexity"` // lets just add validate for now
}