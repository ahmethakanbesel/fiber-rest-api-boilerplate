package controllers

import (
	"fmt"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/app/models"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/pkg/helpers"
	"github.com/ahmethakanbesel/fiber-rest-api-boilerplate/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetFile
// @Summary Get details of the provided file.
// @Tags File
// @Produce json
// @Param file_id   path int true "File ID"
// @Success 200 {object} models.File
// @Router /api/v1/files/{file_id} [get]
func GetFile(ctx *fiber.Ctx) error {
	file := &models.File{}
	err := database.Connection().First(&file, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the file.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "File is fetched.", "data": file})
}

// UploadFile Function for handling file uploads
// @Summary Upload a file to server and return the details.
// @Tags File
// @Accept multipart/form-data
// @Produce json
// @Param file body string true "File"
// @Success 200 {object} models.File
// @Security ApiKeyAuth
// @Router /api/v1/files/ [post]
func UploadFile(c *fiber.Ctx) error {
	claims, err := helpers.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Couldn't get the user details.", "data": err})
	}
	file, err := c.FormFile("file")
	fileUUID := uuid.New()
	// Check for errors:
	if err == nil {
		err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", fileUUID))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't upload the file.", "data:": err})
		}
	}
	newFile := new(models.File)
	newFile.Name = file.Filename
	newFile.MimeType = file.Header.Get("Content-Type")
	newFile.Size = uint32(file.Size)
	newFile.IsPublic = true
	newFile.Path = fileUUID.String()
	newFile.OwnerID = claims.ID
	if err = database.Connection().Create(&newFile).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create a file.", "data:": err})
	}
	return c.JSON(fiber.Map{"status:": "success", "message:": "File is uploaded.", "data": newFile})
}
