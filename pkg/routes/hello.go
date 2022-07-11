package routes

import (
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func HelloRoutes(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
