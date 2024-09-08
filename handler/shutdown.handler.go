package handler

import (
	// "net/http"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func Shutdown(c *fiber.Ctx) error {

	err := exec.Command("shutdown", "/s", "/t", "1").Run()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Shutdown request failed.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"message": "Shutdown request recieved.",
	})
}
