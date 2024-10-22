package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-0815/go-jwt2/controllers"
	initializers "github.com/hiroki-0815/go-jwt2/initializer"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.Run()
}
