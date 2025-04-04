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
	currentUserId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var authorId int
	authorIdStr := c.Query("author_id")
	if authorIdStr == "" {
		authorId = currentUserId
	} else {
		authorId, err = strconv.Atoi(authorIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author_id"})
			return
		}
	}

	posts, err := h.service.Post.GetAllPostsByUserId(authorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch posts"})
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

func (h *Handler) DeletePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = h.service.Post.DeletePost(userId, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) updatePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	var input models.UpdatePostInput
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		return
	}

	if err := h.service.Post.UpdatePost(userId, id, input); err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "ok",
	})
}
