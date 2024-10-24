package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroki-0815/go-jwt2/controllers"
	initializers "github.com/hiroki-0815/go-jwt2/initializer"
	"github.com/hiroki-0815/go-jwt2/middleware"
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
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
