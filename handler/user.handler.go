package handler

import (
	"log"

	"github.com/bebeorca/go-api4/database"
	"github.com/bebeorca/go-api4/models/entity"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerRead(c *fiber.Ctx) error{

	var users []entity.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return c.Status(200).JSON(users)

}