package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable     = "users"
	postsTable     = "posts"
	userPostsTable = "user_posts"
	comments       = "comments"
	postComments   = "post_comments"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// Функция создает подключение к БД
func NewPostgresDB() (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSLMode")))
	if err != nil {
		return nil, err
	}

	// Проверка соединения(отправляет простой запрос к БД)
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
