package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-auth/database"
	"go-auth/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	if err := app.Listen(":8000"); err != nil {
		fmt.Println(err)
		return
	}
}
