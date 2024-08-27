package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB(config *Config) *gorm.DB {
	var dsn string

	if config.Env == "production" {
		dsn = config.DBURL
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database!", err)
	}

	fmt.Println("🧊🧊🧊 Successfully connected to the database!")

	return db
}
