package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go-auth/database"
	"go-auth/routes"
)

func main() {
	err := godotenv.Load() // This will load the .env file
	if err != nil {
		panic("Could not load .env file")
	}

	database.Connect()

	app := fiber.New()

	app.Use(cors.Config{
		AllowCredentials: true,
	})

	routes.Setup(app)

	if err := app.Listen(":8000"); err != nil {
		fmt.Println(err)
		return
	}
}
