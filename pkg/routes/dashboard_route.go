package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// API dashboard

func DashboardRoute(a *fiber.App) {
	a.Get("/dashboard", monitor.New())
}
