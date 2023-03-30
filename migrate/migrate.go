package main

import (
	"web2app/loader"
	"web2app/models"
)

func init() {
	loader.LoadEnvVar()
	loader.ConnectToDB()
}

func main() {
	// Migrate the schema from post struct to database
	loader.DB.AutoMigrate(&models.Post{})
}
