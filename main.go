package main

import (
	"github.com/bebeorca/go-api4/database"
	"github.com/bebeorca/go-api4/database/migrations"
	"github.com/bebeorca/go-api4/routes"
	"github.com/gofiber/fiber/v2"
)

func main(){
	database.DatabaseInit()
	migrations.RunMigration()
	
	app := fiber.New()

	routes.RouteInit(app)


	app.Listen(":8080")
}