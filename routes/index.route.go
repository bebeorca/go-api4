package routes

import (
	"github.com/bebeorca/go-api4/config"
	"github.com/bebeorca/go-api4/handler"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App){
	r.Static("/images", config.ProjectRootPath+"/public/uploads")
	r.Get("/", handler.UserHandlerRead)
	r.Post("/register", handler.UserHandlerCreate)
	r.Post("/upload", handler.UploadPost)

	r.Get("/token", handler.GetToken)
	r.Put("/token/:token", handler.RedeemToken)
}
