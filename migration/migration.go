package migration

import (
	"firstApp/conf"
	"firstApp/models"
)

func Migration() {
	// Migrate the schema
	conf.DB.AutoMigrate(&models.User{})
}
