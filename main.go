package main

import (
	"github.com/gin-gonic/gin"
	initializers "github.com/hiroki-0815/go-jwt2/initializer"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
