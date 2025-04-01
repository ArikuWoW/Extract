package handler

import (
	"fmt"
	"net/http"

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
