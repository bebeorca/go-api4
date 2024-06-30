package migrations

import (
	"log"

	"github.com/bebeorca/go-api4/database"
	"github.com/bebeorca/go-api4/models/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database migrated!")
}