package conf

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Connect to Db
	var err error
	dsn := os.Getenv("DSN")
	dbName := os.Getenv("DB_NAME")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	DB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)

	// Create the database

}
