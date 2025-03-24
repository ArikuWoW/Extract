package models

type Post struct {
	Id        int    `json:"id"`
	Content   string `json:"content" binding:"required"`
	Create_at string `json:"create_at"`
	Like      int    `json:"like"`
	Dislike   int    `json:"dislike"`
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
