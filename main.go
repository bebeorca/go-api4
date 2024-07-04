package main

import (
	"os"
	"log"
	"github.com/bebeorca/go-api4/database"
	"github.com/bebeorca/go-api4/database/migrations"
	"github.com/bebeorca/go-api4/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){

	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowOrigins:     "*",
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	database.DatabaseInit()
	migrations.RunMigration()

	routes.RouteInit(app)
	app.Static("/uploads", "./public/uploads")


	port := os.Getenv("PORT")

	if port == ""{
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
