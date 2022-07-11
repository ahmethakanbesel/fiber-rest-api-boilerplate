package utils

import (
	"fmt"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/pkg/configs"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/pkg/middleware"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreateServer(port int) {
	// Create Fiber App
	config := configs.FiberConfig()
	app := fiber.New(config)

	// Middlewares
	//app.Use(middleware.Example)
	middleware.FiberMiddleware(app)

	// Mount routes
	routes.HelloRoutes(app)
	routes.ApiRoutes(app)
	routes.SwaggerRoute(app)
	routes.DashboardRoute(app)
	routes.NotFoundRoute(app)

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
