package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	dsn := "host=localhost user=levylv password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // dns: data source name
	if err != nil {
		panic("Could not connect to the database")
	}
}
