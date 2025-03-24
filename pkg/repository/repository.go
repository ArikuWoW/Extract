package repository

import (
	"github.com/ArikuWoW/extract/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(login string) (models.User, error)
}

type Post interface {
}

type Comment interface {
}

type Repository struct {
	Authorization
	Post
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
	}
}
