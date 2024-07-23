package handler

import (
	"log"

	"github.com/bebeorca/go-api4/database"
	"github.com/bebeorca/go-api4/models/entity"
	"github.com/bebeorca/go-api4/models/request"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetToken(c *fiber.Ctx) error {
	var token []entity.Token

	uResult := database.DB.Find(&token)

	if uResult.Error != nil {
		log.Fatal(uResult.Error)
	}

	return c.Status(200).JSON(fiber.Map{
		"data": token,
	})
}

func RedeemToken(c *fiber.Ctx) error {

	tokenRequest := new(request.TokenUpdateRequest)

	if err := c.BodyParser(tokenRequest); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	var token entity.Token
	_token := c.Params("token")

	err := database.DB.First(&token, "token = ?", _token).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Token not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	if token.IsRedeemed {
		return c.JSON(fiber.Map{
			"message": "Token sudah digunakan/expired.",
		})
	}else{
		token.IsRedeemed = tokenRequest.IsRedeemed
	}

	errUpdate := database.DB.Save(&token).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": token,
	})

}
