package config

import (
	"log"
	"os"

	"state-tv-api/models" 

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found. Relying on system environment variables.")
	}

	//  Fetch the DSN from the environment
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is missing!")
	}

	// Connect to the database using the hidden string
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection successfully opened")

	// 👇 Told GORM to build your tables automatically 👇
	err = DB.AutoMigrate(&models.Article{}, &models.Subscriber{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	
	log.Println("Database tables successfully migrated!")
}