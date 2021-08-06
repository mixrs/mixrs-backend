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

	channelId := c.Param("channelId")

	createdPost, err := h.PostOperations.CreatePost(&post, channelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdPost})
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

func (h *Handler) GetPostById(c *gin.Context) {
	channelId := c.Param("channelId")
	postId := c.Param("postId")
	post, err := h.PostOperations.GetPostById(channelId, postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}
