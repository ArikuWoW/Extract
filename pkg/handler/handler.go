package handler

import (
	"github.com/ArikuWoW/extract/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/login", h.Login)
	}

	posts := router.Group("/post", h.userIdentity)
	{
		posts.POST("/createPost", h.createPost)
		posts.GET("/getAllPosts", h.getAllPosts)
		posts.GET("/:id", h.getPostById)
		posts.DELETE("/:id", h.DeletePost)
		posts.PUT("/:id", h.updatePost)
	}

	return router
}
