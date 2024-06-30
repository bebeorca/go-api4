package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?ssl-mode=REQUIRED", dbUser, dbPassword, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect to database...")
	}

	fmt.Println("Connecting to database...")
}
