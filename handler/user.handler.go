package handler

import (
	//"fmt"
	"log"
	"time"
	// "time"

	"github.com/bebeorca/go-api4/database"
	"github.com/bebeorca/go-api4/models/entity"
	"github.com/bebeorca/go-api4/models/request"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerRead(c *fiber.Ctx) error{

	var users []entity.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return c.Status(200).JSON(fiber.Map{
		"data": users,
	})

}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := c.BodyParser(user); err != nil{
		return err
	}

	newUser := entity.User{
		Nama: user.Nama,
		Address: user.Address,
		Phone: user.Phone,
		Email: user.Email,
		CreatedAt: time.Now().Local().Format("2006-01-02"),
	}

	errCreating := database.DB.Create(&newUser).Error

	if errCreating != nil{
		return c.Status(500).JSON(fiber.Map{
			"message": errCreating,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Stored!",
		"data": newUser,
	})

}