package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func (app App) GetHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "available",
		"date":   time.Now(),
	})
}
