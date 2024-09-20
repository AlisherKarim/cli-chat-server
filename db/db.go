package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	userStorage
	chatRoomStorage
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgreStorage() (*PostgresStorage, error) {
	var (
		host     = os.Getenv("DATABASE_HOST")
		port     = os.Getenv("DATABASE_PORT")
		user     = os.Getenv("DATABASE_USER")
		password = os.Getenv("DATABASE_PASSWORD")
		dbname   = os.Getenv("DATABASE_NAME")
	)

	var connStr string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	dbObj, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db: dbObj}, nil
}

func (pStorage *PostgresStorage) Init() error {
	return pStorage.createUsersTable()
}

