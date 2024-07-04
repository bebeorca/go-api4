package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/bebeorca/go-api4/database"
	"github.com/bebeorca/go-api4/models/entity"
	"github.com/bebeorca/go-api4/models/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UploadPost(c *fiber.Ctx) error {
	post := new(request.PostRequest)
	if err := c.BodyParser(post); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(post)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failded",
			"error":   errValidate.Error(),
		})
	}

	file, errFile := c.FormFile("image")
	if errFile != nil {
		log.Panic("error file = ", errFile)
	}

	var filename string
	if file != nil {
		filename = file.Filename

		errSaving := c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", filename))
		if errSaving != nil {
			log.Panic("Failed storing image")
		}
	}else{
		c.JSON(fiber.Map{
			"message": "Files cannot be empty.",
		})
	}

	newPost := entity.Post{
		ImagePath: filename,
		CreatedAt: time.Now().Local().Format("2006-01-02"),
	}

	errCreating := database.DB.Create(&newPost).Error
	if errCreating != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": errCreating,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Stored!",
		"data":    newPost,
	})
}
