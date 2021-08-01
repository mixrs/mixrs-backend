package handler

import (
	"bytes"
	"io"
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
	userName := c.PostForm("name")
	userEmail := c.PostForm("email")
	userImage, err := c.FormFile("image")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	file, err := userImage.Open()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.CreateUserDTO{
		Name:  userName,
		Email: userEmail,
		Image: buf.Bytes(),
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
