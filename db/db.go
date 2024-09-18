package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/alisherkarim/cli-chat-server/models"
	_ "github.com/lib/pq"
)

type Storage interface {
	userStorage
	chatRoomStorage
}

type userStorage interface {
	GetUserById(id string) (models.User, error)
	GetUsers() ([]models.User, error)
	CreateUser(name, email string) (models.User, error)
}

type chatRoomStorage interface {
	CreateRoom(name string) (models.ChatRoom, error)
	GetRoom(id string) (models.ChatRoom, error)
	AddUserToRoom(user models.User) error
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

func (pStorage *PostgresStorage) CreateUser(name, email string) (models.User, error) {
	query := `INSERT INTO users (username, email)
						VALUES ($1, $2);`
	res, err := pStorage.db.Exec(query, name, email)
	if err != nil {
		return models.User{}, err
	}

	id, _ := res.LastInsertId()

	return models.User{Username: name, Email: email, Id: int(id)}, nil
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