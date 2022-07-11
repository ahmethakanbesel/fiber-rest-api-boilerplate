package main

import (
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/pkg/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/swagger"
	"os"
	"strconv"

	_ "github.com/ahmethakanbesel/fiber-rest-api-boilerplate/docs"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/platform/migrations"
	_ "github.com/joho/godotenv/autoload"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	if !fiber.IsChild() {
		// Application initialization
		migrations.Migrate()
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	utils.CreateServer(port)
}
