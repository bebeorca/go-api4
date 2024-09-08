package routes

import (
	"github.com/bebeorca/go-api4/config"
	"github.com/bebeorca/go-api4/handler"
	"github.com/bebeorca/go-api4/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App){
	r.Static("/images", config.ProjectRootPath+"/public/uploads")

	r.Post("/login", handler.LoginHandler)

	r.Get("/", middleware.Auth, handler.UserHandlerRead)
	r.Post("/register", handler.UserHandlerCreate)
	r.Post("/upload", handler.UploadPost)

	r.Get("/token", handler.GetToken)
	r.Post("/token/create", handler.CreateToken)
	r.Put("/token/:token", handler.RedeemToken)

	r.Post("/shutdown", handler.Shutdown)
}
