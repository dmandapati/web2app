package loader

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	// https://github.com/joho/godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
