package models

import "time"

type Post struct {
	Id        int       `json:"id" db:"id"`
	Content   string    `json:"content" binding:"required" db:"content"`
	Create_at time.Time `json:"create_at" db:"create_at"`
	Like      int       `json:"like" db:"likes"`
	Dislikes  int       `json:"dislike" db:"dislikes"`
}

type UserPost struct {
	Id        int
	UserId    int
	PostId    int
	CommentId int
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type PostComment struct {
	Id        int
	PostId    int
	CommentId int
	UserId    int
}

type UpdatePostInput struct {
	Content *string `json:"content"`
}
