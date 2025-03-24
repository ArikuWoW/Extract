package handler

import (
	"net/http"

	"github.com/ArikuWoW/extract/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		logrus.Printf("Error in json reading: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		logrus.Errorf("Failed to create user: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type loginInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var input loginInput
	if err := c.BindJSON(&input); err != nil {
		logrus.Errorf("Error:", err)
		return
	}
	h.service.Authorization.GetUser(input.Login, input.Password, c)

}
