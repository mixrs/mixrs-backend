package main

import (
	"log"

	"github.com/Tak1za/mixr/pkg/channel"
	"github.com/Tak1za/mixr/pkg/comment"
	"github.com/Tak1za/mixr/pkg/dbaccess"
	"github.com/Tak1za/mixr/pkg/handler"
	"github.com/Tak1za/mixr/pkg/models"
	"github.com/Tak1za/mixr/pkg/post"
	"github.com/Tak1za/mixr/pkg/user"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func inject(conn *dbaccess.Env) *gin.Engine {
	userOperations := &user.Service{
		Dbo: conn,
	}

	postOperations := &post.Service{
		Dbo: conn,
	}

	channelOperations := &channel.Service{
		Dbo: conn,
	}

	commentOperations := &comment.Service{
		Dbo: conn,
	}

	router := gin.Default()
	router.Use(cors.Default())

	handler.NewHandler(&handler.Config{
		R:                 router,
		UserOperations:    userOperations,
		PostOperations:    postOperations,
		ChannelOperations: channelOperations,
		CommentOperations: commentOperations,
	})

	return router
}

func main() {
	conn, err := dbaccess.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	conn.DB.AutoMigrate(&models.UserModel{})
	conn.DB.AutoMigrate(&models.PostModel{})
	conn.DB.AutoMigrate(&models.ChannelModel{})
	conn.DB.AutoMigrate(&models.CommentModel{})

	r := inject(conn)

	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
