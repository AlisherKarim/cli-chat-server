package db

import "github.com/alisherkarim/cli-chat-server/models"

type chatRoomStorage interface {
	CreateRoom(name string) (models.ChatRoom, error)
	GetRoom(id string) (models.ChatRoom, error)
	AddUserToRoom(user models.User) error
}

func (pStorage *PostgresStorage) CreateRoom(name string) (models.ChatRoom, error) {
	return models.ChatRoom{}, nil
}

func (pStorage *PostgresStorage) GetRoom(id string) (models.ChatRoom, error) {
	return models.ChatRoom{}, nil
}

func (pStorage *PostgresStorage) AddUserToRoom(user models.User) error {
	return nil
}