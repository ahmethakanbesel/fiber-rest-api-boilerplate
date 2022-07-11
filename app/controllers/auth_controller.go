package controllers

import (
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/app/dto"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/app/models"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

// Login Function for creating JWT
// @Summary If login credentials are correct, returns JWT.
// @Tags Auth
// @Accept json
// @Produce json
// @Param identity body string true "Identity (email)"
// @Param password body string true "Password"
// @Param role body int true "Role ID"
// @Success 200 {object} models.User "If login is successful returns JWT and user details."
// @Router /api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	input := new(dto.AuthDTO)
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid data given."})
	}
	identity := input.Identity
	pass := input.Password
	user := &models.User{}
	err := database.Connection().Joins("UserRole").Joins("Photo").First(&user, "email = ?", identity).Error
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid email or password"})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if user.Email != identity || err != nil {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid email or password"})
	}
	if user.UserRole.Name != input.Role {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid email or password"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["identity"] = identity
	claims["expires"] = time.Now().Add(time.Hour * 72).Unix()
	claims["role"] = user.UserRole.Name
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't sign token", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": user, "token": t})
}
