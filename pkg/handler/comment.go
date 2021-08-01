package handler

import (
	"log"
	"net/http"

	"github.com/Tak1za/mixr/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	var comment models.CreateCommentDTO

	if err := c.ShouldBindJSON(&comment); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId := c.Param("postId")

	commentId, err := h.CommentOperations.CreateComment(&comment, postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": commentId})
}

func (h *Handler) GetComments(c *gin.Context) {
	postId := c.Param("postId")
	comments, err := h.CommentOperations.GetComments(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}
