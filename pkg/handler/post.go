package handler

import (
	"log"
	"net/http"

	"github.com/Tak1za/mixr/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var post models.CreatePostDTO

	if err := c.ShouldBindJSON(&post); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId, err := h.PostOperations.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": postId})
}

func (h *Handler) GetPosts(c *gin.Context) {
	channelId := c.Param("channelId")
	posts, err := h.PostOperations.GetPosts(channelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": posts})
}
