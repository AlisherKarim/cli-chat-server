package db

import "github.com/alisherkarim/cli-chat-server/models"

type roomUsersStorage interface {
	CreateLink(room_id, user_id string) (models.ChatRoom, error)
}

func (pStorage *PostgresStorage) createRoomUsersTable() error {
	return nil
}

// CreateLink implements Storage.
func (pStorage *PostgresStorage) CreateLink(room_id string, user_id string) (models.ChatRoom, error) {
	panic("unimplemented")
}