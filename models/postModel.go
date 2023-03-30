package models

import "gorm.io/gorm"

// create the schema
// // https://gorm.io/docs/models.html#Embedded-Struct
type Post struct {
	gorm.Model
	Title string
	Body  string
}
