package controllers

import (
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/app/dto"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/app/models"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/pkg/helpers"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

// CreateUser
// @Summary Create a new user.
// @Tags User
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} models.User
// @Router /api/v1/users [post]
func CreateUser(ctx *fiber.Ctx) error {
	db := database.Connection()
	newUser := new(dto.UserRegisterDTO)
	if err := ctx.BodyParser(newUser); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid data given.", "data": err})
	}
	user := new(models.User)
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = models.HashPassword(newUser.Password)
	user.UserRoleID = 3
	if err := db.Create(&user).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["identity"] = user.Email
	claims["expires"] = time.Now().Add(time.Hour * 72).Unix()
	claims["role"] = user.UserRoleID
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't sign token", "data": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "User created.", "data": user, "token": t})
}

// GetUsers
// @Summary Get list of all users.
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Security ApiKeyAuth
// @Router /api/v1/users [get]
func GetUsers(ctx *fiber.Ctx) error {
	var users []models.User
	database.Connection().Joins("UserRole").Joins("Photo").Find(&users)
	return ctx.JSON(fiber.Map{"status": "success", "message": "Users are fetched.", "data": users})
}

// GetUser
// @Summary Get details of the current user.
// @Tags User
// @Produce json
// @Param user_id   path int true "User ID"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /api/v1/users/{user_id} [get]
func GetUser(ctx *fiber.Ctx) error {
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't get user", "data": err})
	}
	if claims.Role == "customer" && ctx.Params("id") != strconv.Itoa(int(claims.ID)) {
		return ctx.Status(403).JSON(fiber.Map{"status": "error", "message": "You can't access the user."})
	}
	user := &models.User{}
	err = database.Connection().First(&user, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the user.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": user})
}

// UpdateUser
// @Summary Update details of the current user.
// @Tags User
// @Produce json
// @Accept json
// @Param  User body models.User true "User"
// @Param user_id   path int true "User ID"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /api/v1/users/{user_id} [post]
func UpdateUser(ctx *fiber.Ctx) error {
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't get user", "data": err})
	}
	newUser := new(dto.UserDTO)
	if err = ctx.BodyParser(newUser); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid data given.", "data": err})
	}
	user := &models.User{}
	err = database.Connection().First(&user, "id = ?", claims.ID).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the user.", "data:": err})
	}
	if newUser.Name != "" {
		user.Name = newUser.Name
	}
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Phone != "" {
		user.Phone = newUser.Phone
	}
	if newUser.PhotoID > 0 {
		user.PhotoID = newUser.PhotoID
	}
	if newUser.Password != "" {
		user.Password = models.HashPassword(newUser.Password)
	}
	database.Connection().Save(user)
	return ctx.JSON(fiber.Map{"status": "success", "message": "User updated.", "data": user})
}

// DeleteUser
// @Summary Delete the user.
// @Tags User
// @Produce json
// @Param user_id   path int true "User ID"
// @Success 200 {string} status "OK"
// @Security ApiKeyAuth
// @Router /api/v1/users/{user_id} [delete]
func DeleteUser(ctx *fiber.Ctx) error {
	user := &models.User{}
	err := database.Connection().Delete(&user, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't delete the user.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "User deleted."})
}
