package migration

import (
	"firstApp/conf"
	"firstApp/models"
)

func Migration() {
	// Migrate schemas to database
	models := []interface{}{&models.Task{}, &models.User{}}
	conf.DB.AutoMigrate(models...)
}
