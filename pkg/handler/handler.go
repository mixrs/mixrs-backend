package handler

import (
	"github.com/Tak1za/mixr/pkg/channel"
	"github.com/Tak1za/mixr/pkg/post"
	"github.com/Tak1za/mixr/pkg/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserOperations    user.Operations
	PostOperations    post.Operations
	ChannelOperations channel.Operations
}

type Config struct {
	R                 *gin.Engine
	UserOperations    user.Operations
	PostOperations    post.Operations
	ChannelOperations channel.Operations
}

func NewHandler(c *Config) {
	h := &Handler{
		UserOperations:    c.UserOperations,
		PostOperations:    c.PostOperations,
		ChannelOperations: c.ChannelOperations,
	}

	g := c.R.Group("/api/v1")

	g.GET("/users/:userId", h.GetUser)
	g.POST("/users", h.CreateUser)
	g.PUT("/users/:userId", h.UpdateUser)
	g.DELETE("/users/:userId", h.DeleteUser)

	g.POST("/posts", h.CreatePost)
	g.GET("/posts/channels/:channelId", h.GetPosts)

	g.POST("/channels", h.CreateChannel)
	g.GET("/channels/:channelId", h.GetChannel)
	g.GET("/channels", h.GetChannels)
}
