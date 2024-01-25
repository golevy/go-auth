package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"jwt/database"
	"jwt/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
