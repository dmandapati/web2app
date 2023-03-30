package main

import (
	"web2app/controllers"
	"web2app/loader"

	"github.com/gin-gonic/gin"
)

func init() {
	loader.LoadEnvVar()
	loader.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)

	r.Run()
}
