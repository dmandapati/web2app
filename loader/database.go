package loader

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	log.Println("Database Connected")
	if err != nil {
		log.Fatal("Failed to connect Database")
	}
}
