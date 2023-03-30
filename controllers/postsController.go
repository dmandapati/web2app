package controllers

import (
	"web2app/loader"
	"web2app/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get the data off req body
	var Hello struct {
		Title string
		Body  string
	}
	c.Bind(&Hello)

	// Create post
	post := models.Post{Title: Hello.Title, Body: Hello.Body}

	result := loader.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"Post": post,
	})
}
