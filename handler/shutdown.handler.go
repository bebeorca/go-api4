package handler

import "github.com/gofiber/fiber/v2"

func Shutdown(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"message": "Shutdown request recieved.",
	})
}
