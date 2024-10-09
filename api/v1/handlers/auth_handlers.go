package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alisherkarim/cli-chat-server/api/v1/types"
	"github.com/alisherkarim/cli-chat-server/models"
	"github.com/alisherkarim/cli-chat-server/pkg/response"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwt_secret = os.Getenv("JWT_SECRET")

func (mainHandler *MainHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := &types.LoginRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := mainHandler.userController.GetUserByUsername(req.Username)

	if err != nil {
		response.RespondWithError(w, http.StatusUnauthorized, err)
		return
	}

	expiresAt := time.Now().Add(time.Hour).Unix()
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		resp := "Invalid login credentials. Please try again"
		response.RespondWithErrorMsg(w, http.StatusUnauthorized, resp)
		return
	}

	tk := &models.Token{
		UserId: string(user.Id),
		Username:   user.Username,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		log.Println(err.Error())
		resp := "Something went wrong on out end. Please try again"
		response.RespondWithErrorMsg(w, http.StatusInternalServerError, resp)
		return;
	}

	response.RespondWithJson(w, http.StatusOK, types.LoginResponseBody{User: user, AccessToken: tokenString})
}

func (mainHandler *MainHandler) Register(w http.ResponseWriter, r *http.Request) {
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
	response.RespondWithJson(w, http.StatusCreated, user)
}