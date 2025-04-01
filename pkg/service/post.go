package service

import (
	"github.com/ArikuWoW/extract/models"
	"github.com/ArikuWoW/extract/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(userId int, post models.Post) (int, error) {
	return s.repo.CreatePost(userId, post)
}

func (s *PostService) GetAllPostsByUserId(userId int) ([]models.Post, error) {
	return s.repo.GetAllPostsByUserId(userId)
}

func (s *PostService) GetPostById(userId, postId int) (models.Post, error) {
	return s.repo.GetPostById(userId, postId)
}
