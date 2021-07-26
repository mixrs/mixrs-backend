package handler

import (
	"log"
	"net/http"

	"github.com/Tak1za/mixr/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	userId := c.Param("userId")
	fetchedUser, err := h.UserOperations.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": fetchedUser})
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.CreateUserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := h.UserOperations.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": userId})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	userId := c.Param("userId")
	var updateUser models.UpdateUserDTO

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserOperations.UpdateUser(&updateUser, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	if err := h.UserOperations.DeleteUser(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
