package controllers

import (
	"web2app/loader"
	"web2app/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get the data off req body
	// https://gorm.io/docs/models.html#Embedded-Struct
	var Hello struct {
		Title string
		Body  string
	}
	c.Bind(&Hello)

	// Create post
	// https://gorm.io/docs/create.html#Create-Record
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

func PostsIndex(c *gin.Context) {
	// https://gorm.io/docs/query.html#Retrieving-all-objects
	// Get the post
	var posts []models.Post
	loader.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"Posts": posts,
	})
}
