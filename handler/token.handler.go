package handler

import (
	"log"

	"github.com/bebeorca/go-api4/database"
	"github.com/bebeorca/go-api4/models/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateToken(c *fiber.Ctx) error {
	user := new(entity.Token)
	if err := c.BodyParser(user); err != nil{
		return err
	}

	newToken := entity.Token{
		Token: user.Token,
	}

	errCreating := database.DB.Create(&newToken).Error

	if errCreating != nil{
		return c.Status(500).JSON(fiber.Map{
			"message": errCreating,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Stored!",
		"data": newToken,
	}) 
}

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
		return c.Status(409).JSON(fiber.Map{
			"message": "Token sudah digunakan/expired.",
		})
	}else{
		token.IsRedeemed = true
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
