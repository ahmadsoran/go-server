package conf

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func EnvLoader() {
	// if production use this
	gin.SetMode(gin.DebugMode)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
