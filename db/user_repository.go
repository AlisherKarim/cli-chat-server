package db

import (
	"database/sql"
	"fmt"

	"github.com/alisherkarim/cli-chat-server/models"
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
    user_id SERIAL PRIMARY KEY,
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
	query := `SELECT user_id, username, email FROM users WHERE username = $1;`
	row := pStorage.db.QueryRow(query, username)
	user := models.User{}
	if err := row.Scan(&user.Id, &user.Username, &user.Email); err != nil {
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
		return []models.User{}, err
	}
	defer rows.Close()
	users := []models.User{}
	
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (pStorage *PostgresStorage) CreateUser(name, email, password_hash string) (models.User, error) {
	query := `INSERT INTO users (username, email, password)
						VALUES ($1, $2, $3);`

	_, err := pStorage.db.Exec(query, name, email, password_hash)
	if err != nil {
		return models.User{}, err
	}

	return models.User{Username: name, Email: email}, nil
}