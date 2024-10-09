package db

import (
	"database/sql"
	"fmt"

	"github.com/alisherkarim/cli-chat-server/models"
	"github.com/google/uuid"
)

type userStorage interface {
	GetUserById(id string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetUsers() ([]models.User, error)
	CreateUser(name, email, password_hash string) (models.User, error)
}

func (pStorage *PostgresStorage) createUsersTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
    user_id VARCHAR(255) UNIQUE NOT NULL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`
	_, err := pStorage.db.Exec(query)
	
	if err != nil {
		return err
	}

	return nil
}

func (pStorage *PostgresStorage) GetUserById(id string) (models.User, error) {
	query := `SELECT user_id, username, email FROM users WHERE user_id = $1;`
	row := pStorage.db.QueryRow(query, id)
	user := models.User{}
	if err := row.Scan(&user.Id, &user.Username, &user.Email); err != nil {
		if err == sql.ErrNoRows {
				return models.User{}, fmt.Errorf("user with id %s not found", id)
		}
		return models.User{}, err
	}

	return user, nil
}

func (pStorage *PostgresStorage) GetUserByUsername(username string) (models.User, error) {
	query := `SELECT user_id, username, email, password FROM users WHERE username = $1;`
	row := pStorage.db.QueryRow(query, username)
	user := models.User{}
	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
				return models.User{}, fmt.Errorf("user with id %s not found", username)
		}
		return models.User{}, err
	}

	return user, nil
}

func (pStorage *PostgresStorage) GetUsers() ([]models.User, error) {
	query := `SELECT user_id, username, email FROM users;`
	rows, err := pStorage.db.Query(query)
	if err != nil {
			return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
			user := models.User{}
			err := rows.Scan(&user.Id, &user.Username, &user.Email)
			if err != nil {
					return nil, fmt.Errorf("failed to scan user: %w", err)
			}
			users = append(users, user)
	}

	return users, nil
}

func (pStorage *PostgresStorage) CreateUser(name, email, passwordHash string) (models.User, error) {
    userId := uuid.New().String()
    query := `INSERT INTO users (user_id, username, email, password) VALUES ($1, $2, $3, $4);`

    _, err := pStorage.db.Exec(query, userId, name, email, passwordHash)
    if err != nil {
        return models.User{}, fmt.Errorf("failed to create user: %w", err)
    }

    return models.User{Id: userId, Username: name, Email: email}, nil
}
