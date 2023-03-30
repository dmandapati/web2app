package main

import (
	"web2app/loader"

	"github.com/gin-gonic/gin"
)

func init() {
	loader.LoadEnvVar()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
