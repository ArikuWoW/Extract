package service

import (
	"github.com/ArikuWoW/extract/models"
	"github.com/ArikuWoW/extract/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(login, password string, c *gin.Context)
	ParseToken(accessToken string) (int, error)
}

type Post interface {
	CreatePost(userId int, post models.Post) (int, error)
}

type Comment interface {
}

type Service struct {
	Authorization
	Post
	Comment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Post:          NewPostService(repos.Post),
	}
}
