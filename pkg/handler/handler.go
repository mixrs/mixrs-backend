package handler

import (
	"github.com/Tak1za/mixr/pkg/channel"
	"github.com/Tak1za/mixr/pkg/comment"
	"github.com/Tak1za/mixr/pkg/post"
	"github.com/Tak1za/mixr/pkg/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserOperations    user.Operations
	PostOperations    post.Operations
	ChannelOperations channel.Operations
	CommentOperations comment.Operations
}

type Config struct {
	R                 *gin.Engine
	UserOperations    user.Operations
	PostOperations    post.Operations
	ChannelOperations channel.Operations
	CommentOperations comment.Operations
}

func NewHandler(c *Config) {
	h := &Handler{
		UserOperations:    c.UserOperations,
		PostOperations:    c.PostOperations,
		ChannelOperations: c.ChannelOperations,
		CommentOperations: c.CommentOperations,
	}

	g := c.R.Group("/api/v1")

	g.GET("/users/:userId", h.GetUser)
	g.POST("/users", h.CreateUser)
	g.PUT("/users/:userId", h.UpdateUser)
	g.DELETE("/users/:userId", h.DeleteUser)

	g.POST("/channels", h.CreateChannel)
	g.GET("/channels/:channelId", h.GetChannel)
	g.GET("/channels", h.GetChannels)

	g.POST("/channels/:channelId/posts", h.CreatePost)
	g.GET("/channels/:channelId/posts", h.GetPosts)
	g.GET("/channels/:channelId/posts/:postId", h.GetPostById)

	g.POST("/channels/:channelId/posts/:postId/comments", h.CreateComment)
	g.GET("/channels/:channelId/posts/:postId/comments", h.GetComments)
}
