package main

import (
	"io"

	"github.com/evansopilo/visuai/pkg/data"
	"github.com/gofiber/fiber/v2"
)

func (app App) UploadFile(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
		})
	}

	buffer, err := file.Open()
	if err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
		})
	}
	defer buffer.Close()

	fileByte, err := io.ReadAll(buffer)
	if err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
		})
	}

	url, err := app.BlobModel.UploadBytesToBlob(fileByte, map[string]string{})
	if err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
		})
	}

	if err := app.Models.Post.UpdateByID(c.Context(), c.FormValue("post_id"), &data.Post{PhotoURL: url}); err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": map[string]string{
			"id": c.FormValue("post_id"),
		},
	})
}
