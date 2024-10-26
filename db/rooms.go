package db

import (
	"database/sql"
	"fmt"

	"github.com/alisherkarim/cli-chat-server/models"
	"github.com/google/uuid"
)

type chatRoomStorage interface {
	CreateRoom(name string) (models.ChatRoom, error)
	GetRoom(id string) (models.ChatRoom, error)
	GetRooms() ([]models.ChatRoom, error)
}


// NOTE: temporary code. removed after db migrations are implemented
func (pStorage *PostgresStorage) createRoomTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS rooms (
    room_id VARCHAR(255) UNIQUE NOT NULL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := pStorage.db.Exec(query)
	
	if err != nil {
		return err
	}

	return nil
}

func (pStorage *PostgresStorage) CreateRoom(name string) (models.ChatRoom, error) {
	roomId := uuid.New().String()
	query := `INSERT INTO rooms (room_id, name) VALUES ($1, $2);`

	_, err := pStorage.db.Exec(query, roomId, name)

	if err != nil {
		return models.ChatRoom{}, fmt.Errorf("failed to create room with name %s: %w", name, err)
	}

	return models.ChatRoom{Id: roomId, Name: name}, nil
}

func (pStorage *PostgresStorage) GetRoom(id string) (models.ChatRoom, error) {
	query := `SELECT room_id, name FROM rooms WHERE room_id = $1;`
	row := pStorage.db.QueryRow(query, id)
	room := models.ChatRoom{}
	if err := row.Scan(&room.Id, &room.Name); err != nil {
		if err == sql.ErrNoRows {
			return models.ChatRoom{}, fmt.Errorf("room with id %s not found", id)
		}
		return models.ChatRoom{}, err
	}

	return models.ChatRoom{}, nil
}

func (pStorage *PostgresStorage) GetRooms() ([]models.ChatRoom, error) {
	return []models.ChatRoom{}, nil
}
