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
	createPostQuery := fmt.Sprintf("INSERT INTO %s (content, create_at) VALUES ($1, $2) RETURNING id", postsTable)
	row := tx.QueryRow(createPostQuery, post.Content, now)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUSersPostQuery := fmt.Sprintf("INSERT INTO %s (user_id, post_id) VALUES ($1, $2)", userPostsTable)
	_, err = tx.Exec(createUSersPostQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *PostDB) GetAllPostsByUserId(userId int) ([]models.Post, error) {
	var posts []models.Post
	query := fmt.Sprintf("SELECT p.id, p.content, p.create_at, p.likes, p.dislikes FROM %s p INNER JOIN %s u on p.id = u.post_id WHERE u.user_id = $1", postsTable, userPostsTable)
	err := r.db.Select(&posts, query, userId)

	return posts, err
}

func (r *PostDB) GetPostById(userId, postId int) (models.Post, error) {
	var post models.Post

	query := fmt.Sprintf("SELECT p.id, p.id, p.content, p.create_at, p.likes, p.dislikes FROM %s p INNER JOIN %s u on p.id = u.post_id WHERE u.user_id = $1 AND u.post_id = $2", postsTable, userPostsTable)
	err := r.db.Get(&post, query, userId, postId)

	return post, err
}

func (r *PostDB) DeletePost(userId, postId int) error {
	query := fmt.Sprintf("DELETE FROM %s p USING %s u WHERE p.id = u.post_id AND u.user_id=$1 AND u.post_id=$2", postsTable, userPostsTable)

	_, err := r.db.Exec(query, userId, postId)
	return err
}
