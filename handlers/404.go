package handlers

import "github.com/gofiber/fiber/v2"

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{
		"success": false,
	})
}
