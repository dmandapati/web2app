package controllers

import (
	"log"
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

	// Return it
	c.JSON(200, gin.H{
		"Posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// https://gorm.io/docs/query.html#Retrieving-objects-with-primary-key
	// Get ID from url
	id := c.Param("id")

	// Find The Post
	var post models.Post
	loader.DB.First(&post, id)

	// Return it
	c.JSON(200, gin.H{
		"Post": post,
	})
}

func PostUpdate(c *gin.Context) {

	// Get the post ID
	id := c.Param("id")

	// Get the data of requested post
	var Hello struct {
		Title string
		Body  string
	}
	c.Bind(&Hello)

	// find the post
	var post models.Post
	loader.DB.First(&post, id)

	// update it
	// https://gorm.io/docs/update.html#Updates-multiple-columns
	loader.DB.Model(&post).Updates(models.Post{
		Title: Hello.Title,
		Body:  Hello.Body,
	})

	// Return it
	c.JSON(200, gin.H{
		"Post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get the POST id
	id := c.Param("id")

	// delete the post
	// https://gorm.io/docs/delete.html#Delete-with-primary-key
	loader.DB.Delete(&models.Post{}, id)

	// Return it
	c.Status(200)
	log.Println("Record deleted")
}
