package database

import (
	"fmt"
	"go-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB // DB is a global variable to hold the database connection.

func Connect() {
	// Use os.Getenv to read the environment variables
	// Use os.Getenv to read the environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"))

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // dns: data source name

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

	// AutoMigrate ensures the database schema matches the User model structure.
	// This creates or updates the User table in the database based on the User struct definition in models package.
	err = connection.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		return
	}
}
