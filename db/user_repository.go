package db

import (
	"database/sql"
	"fmt"

	"github.com/alisherkarim/cli-chat-server/models"
)

type userStorage interface {
	GetUserById(id string) (models.User, error)
	GetUsers() ([]models.User, error)
	CreateUser(name, email, password_hash string) (models.User, error)
}

func (pStorage *PostgresStorage) createUsersTable() error {
	query := `create table if not exists users (
		USER_ID SERIAL PRIMARY KEY,
		USERNAME VARCHAR(50) UNIQUE NOT NULL,
		EMAIL VARCHAR(50) UNIQUE NOT NULL
	)`
	_, err := pStorage.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (pStorage *PostgresStorage) GetUserById(id string) (models.User, error) {
	query := `SELECT USER_ID, USERNAME, EMAIL FROM users WHERE USER_ID = $1;`
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

func (pStorage *PostgresStorage) GetUsers() ([]models.User, error) {
	query := `SELECT USER_ID, USERNAME, EMAIL FROM users;`
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
	query := `INSERT INTO users (username, email)
						VALUES ($1, $2);`
	_, err := pStorage.db.Exec(query, name, email)
	if err != nil {
		return models.User{}, err
	}

	return models.User{Username: name, Email: email}, nil
}