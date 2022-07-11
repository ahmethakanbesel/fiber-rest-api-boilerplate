package routes

import (
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/app/controllers"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// ApiRoutes func for describe group of public routes.
func ApiRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Example restricted route
	route.Get("/restricted", middleware.Protected(), controllers.Hello)

	// Routes for auth
	route.Post("/auth/login", controllers.Login)

	// User routes
	route.Get("/users", controllers.GetUsers)
	route.Get("/users/:id", middleware.Protected(), controllers.GetUser)
	route.Post("/users/", controllers.CreateUser)
	route.Post("/users/:id", middleware.Protected(), controllers.UpdateUser)
	route.Delete("/users/:id", middleware.Protected(), controllers.DeleteUser)

	// File Routes
	route.Get("/files/:id", controllers.GetFile)
	route.Post("/files/", middleware.Protected(), controllers.UploadFile)
}
