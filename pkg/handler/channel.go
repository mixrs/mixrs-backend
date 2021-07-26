package handler

import (
	"log"
	"net/http"

	"github.com/Tak1za/mixr/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateChannel(c *gin.Context) {
	var channel models.CreateChannelDTO

	if err := c.ShouldBindJSON(&channel); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	channelId, err := h.ChannelOperations.CreateChannel(&channel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": channelId})
}

func (h *Handler) GetChannel(c *gin.Context) {
	channelId := c.Param("channelId")
	fetchedChannel, err := h.ChannelOperations.GetChannel(channelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": fetchedChannel})
}

func (h *Handler) GetChannels(c *gin.Context) {
	channels, err := h.ChannelOperations.GetChannels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": channels})
}
