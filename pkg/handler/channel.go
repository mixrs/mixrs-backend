package handler

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/Tak1za/mixr/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateChannel(c *gin.Context) {
	channelTitle := c.PostForm("title")
	channelDescription := c.PostForm("description")
	channelImage, err := c.FormFile("avatar")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := channelImage.Open()
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

	channel := models.CreateChannelDTO{
		Title:       channelTitle,
		Description: channelDescription,
		Image:       buf.Bytes(),
	}

	createdChannel, err := h.ChannelOperations.CreateChannel(&channel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdChannel})
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
