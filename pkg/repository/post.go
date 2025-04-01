package repository

import (
	"fmt"
	"time"

	"github.com/ArikuWoW/extract/models"
	"github.com/jmoiron/sqlx"
)

type PostDB struct {
	db *sqlx.DB
}

func NewPostDB(db sqlx.DB) *PostDB {
	return &PostDB{db: &db}
}

func (r *PostDB) CreatePost(userId int, post models.Post) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	now := time.Now()
	createPostQuery := fmt.Sprintf("INSERT INTO %s (content, create_at) VALUES ($1, $2) RETURNING id", posts)
	row := tx.QueryRow(createPostQuery, post.Content, now)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUSersPostQuery := fmt.Sprintf("INSERT INTO %s (user_id, post_id) VALUES ($1, $2)", userPosts)
	_, err = tx.Exec(createUSersPostQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
