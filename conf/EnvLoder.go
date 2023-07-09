package conf

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvLoader() {
	// if production use this
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
