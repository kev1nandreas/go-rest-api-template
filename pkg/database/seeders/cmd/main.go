package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kev1nandreas/go-rest-api-template/env"
	"github.com/kev1nandreas/go-rest-api-template/pkg/database/seeders"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a seeder action: 'up' or 'down'")
	}

	action := os.Args[1]

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	db := connectDatabase()

	switch action {
	case "up":
		log.Println("Running database seeders (up)...")
		if err := seeders.RunAllSeeders(db); err != nil {
			log.Fatalf("Failed to run seeders: %v", err)
		}
		log.Println("Seeders completed successfully!")
	case "down":
		log.Println("Clearing all seeded data (down)...")
		if err := seeders.ClearAllData(db); err != nil {
			log.Fatalf("Failed to clear data: %v", err)
		}
		log.Println("Data cleared successfully!")
	default:
		log.Fatalf("Unknown action: %s. Use 'up' or 'down'", action)
	}
}

func connectDatabase() *gorm.DB {
	var database *gorm.DB
	var err error

	db_hostname := env.GetEnvString("POSTGRES_HOST", "localhost")
	db_name := env.GetEnvString("POSTGRES_DB", "postgres")
	db_user := env.GetEnvString("POSTGRES_USER", "postgres")
	db_pass := env.GetEnvString("POSTGRES_PASSWORD", "password")
	db_port := env.GetEnvString("POSTGRES_PORT", "5432")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		db_hostname,
		db_user,
		db_pass,
		db_name,
		db_port,
	)

	for i := 1; i <= 3; i++ {
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		} else {
			log.Printf("Attempt %d: Failed to connect to database. Retrying...", i)
			time.Sleep(3 * time.Second)
		}
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after 3 attempts: %v", err)
	}

	log.Println("Successfully connected to database")
	return database
}
