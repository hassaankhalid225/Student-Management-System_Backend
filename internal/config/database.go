package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(config Config) *gorm.DB {
	var db *gorm.DB
	var err error

	// Check if we should use SQLite (for development)
	if config.DBHost == "sqlite" {
		log.Println("Using SQLite database for development")
		db, err = gorm.Open(sqlite.Open("sms.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		// Use PostgreSQL (for production)
		log.Println("Using PostgreSQL database")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db
	return db
}
