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
	database.DatabaseInit()
	migrations.RunMigration()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowOrigins: "https://vue-app-cynu.vercel.app/", // Allow all origins
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

	routes.RouteInit(app)


	port := os.Getenv("PORT")

	if port == ""{
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}