package main

import (
	"gin/server/configs"
	"gin/server/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	configs.ConnectDB()
	routes.LetterRoute(r)
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "Hello from gin-gonic & mongo",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
