package handler

import (
	//"fmt"
	"log"
	"time"

	// "time"

	"github.com/bebeorca/go-api4/database"
	"github.com/bebeorca/go-api4/models/entity"
	"github.com/bebeorca/go-api4/models/request"
	"github.com/bebeorca/go-api4/utils"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerRead(c *fiber.Ctx) error {

	var users []entity.User
	var posts []entity.Post

	uresult := database.DB.Find(&users)
	presult := database.DB.Find(&posts)

	if uresult.Error != nil {
		log.Fatal(uresult.Error)
	}

	if presult.Error != nil {
		log.Fatal(presult.Error)
	}

	return c.Status(200).JSON(fiber.Map{
		"data":  users,
		"image": posts,
	})

}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	newUser := entity.User{
		Nama:      user.Nama,
		Address:   user.Address,
		Phone:     user.Phone,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now().Local().Format("2006-01-02"),
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreating := database.DB.Create(&newUser).Error

	if errCreating != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": errCreating,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Stored!",
		"data":    newUser,
	})

}
