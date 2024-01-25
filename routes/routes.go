package routes

import (
	"github.com/gofiber/fiber/v2"
	"jwt/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
