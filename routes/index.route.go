package routes

import (
	"github.com/bebeorca/go-api4/handler"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App){
	api := r.Group("/api")
	api.Get("/", handler.UserHandlerRead)
	api.Post("/register", handler.UserHandlerCreate)
}
