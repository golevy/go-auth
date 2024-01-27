package database

import (
	"fmt"
	"go-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	dsn := "host=localhost user=levylv password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // dns: data source name

	if err != nil {
		panic("Could not connect to the database")
	}

	// AutoMigrate ensures the database schema matches the User model structure.
	// This creates or updates the User table in the database based on the User struct definition in models package.
	err = connection.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		return
	}
}
