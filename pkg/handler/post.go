package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ArikuWoW/extract/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input models.Post
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		return
	}

	id, err := h.service.Post.CreatePost(userId, input)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllPostsResp struct {
	Data []models.Post `json:"data"`
}

func (h *Handler) getAllPosts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	posts, err := h.service.Post.GetAllPostsByUserId(userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, getAllPostsResp{
		Data: posts,
	})
}

func (h *Handler) getPostById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	post, err := h.service.Post.GetPostById(userId, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, post)
}
