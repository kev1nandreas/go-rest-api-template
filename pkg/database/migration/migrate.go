package migration

import (
	"github.com/kev1nandreas/go-rest-api-template/pkg/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Enable UUID extension for PostgreSQL
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return err
	}

	return db.AutoMigrate(&models.Book{}, &models.User{})
}