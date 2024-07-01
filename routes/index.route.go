package routes

import (
	"github.com/bebeorca/go-api4/handler"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App){
	r.Get("/", handler.UserHandlerRead)
	r.Post("/register", handler.UserHandlerCreate)
}
