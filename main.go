package main

import (
	"github.com/gin-gonic/gin"
	"go-socket/modules/message/handlers"
	"go-socket/modules/message/repositories"
)

var (
	handler *handlers.MessageHandler
)

func init() {
	stg := repositories.NewMessageStorage()
	handler = handlers.NewMessageHandler(stg)
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		api.POST("/message", handler.SubmitMessage)
		api.GET("/messages", handler.GetAllMessages)
	}

	return r
}
