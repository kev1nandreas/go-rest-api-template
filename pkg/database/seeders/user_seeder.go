package seeders

import (
	"log"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/kev1nandreas/go-rest-api-template/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB, count int) error {
	gofakeit.Seed(0) // Use 0 for random seed or set a fixed number for reproducible data

	for i := 0; i < count; i++ {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := models.User{
			ID:       uuid.New(),
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		if err := db.Create(&user).Error; err != nil {
			log.Printf("Failed to create user: %v", err)
			return err
		}
	}

	log.Printf("Successfully seeded %d users", count)
	return nil
}

func ClearUsers(db *gorm.DB) error {
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		return err
	}
	log.Println("Successfully cleared users table")
	return nil
}
