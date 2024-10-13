package controllers

import (
	"github.com/alisherkarim/cli-chat-server/db"
	"github.com/alisherkarim/cli-chat-server/models"
)


type UserController struct {
	storage db.Storage
}

func NewUserController(storage db.Storage) *UserController {
	return &UserController{storage: storage}
}

func (userController *UserController) CreateUser(name, email, password_hash string) (models.User, error) {
	user, err := userController.storage.CreateUser(name, email, password_hash)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (userController *UserController) GetUsers() ([]models.User, error) {
	users, err := userController.storage.GetUsers()
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func (userController *UserController) GetUserById(id string) (models.User, error) {
	user, err := userController.storage.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (userController *UserController) GetUserByUsername(username string) (models.User, error) {
	user, err := userController.storage.GetUserByUsername(username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}