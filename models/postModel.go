package models

import "gorm.io/gorm"

// create the schema
type Post struct {
	gorm.Model
	Title string
	Body  string
}
