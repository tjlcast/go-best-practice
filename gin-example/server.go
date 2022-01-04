package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(context *gin.Context) {
		context.String(200, "%s", "pong!")
	})
	log.Fatalln(server.Run("localhost:8888"))
}