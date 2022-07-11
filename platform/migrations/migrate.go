package migrations

import (
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/app/models"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/platform/database"
)

func Migrate() {
	database.Connection().AutoMigrate(&models.File{})
	database.Connection().AutoMigrate(&models.Permission{})
	database.Connection().AutoMigrate(&models.RolePermission{})
	database.Connection().AutoMigrate(&models.User{})
	database.Connection().AutoMigrate(&models.UserRole{})
}
