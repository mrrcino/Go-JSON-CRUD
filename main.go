package main

import (
	"Go-CRUD/controllers"
	"Go-CRUD/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.PUT("/posts:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts:id", controllers.PostShow)
	r.DELETE("/posts:id", controllers.PostDelete)

	r.Run()
}
