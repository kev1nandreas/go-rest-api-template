package database

import (
	"fmt"
	"log"
	"time"

	"github.com/kev1nandreas/go-rest-api-template/env"
	"github.com/kev1nandreas/go-rest-api-template/pkg/database/migration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Offset(offset int) *gorm.DB
	Limit(limit int) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) Database
	Delete(interface{}, ...interface{}) *gorm.DB
	Model(model interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) Database
	Updates(interface{}) *gorm.DB
	Order(value interface{}) *gorm.DB
	Error() error
}

type GormDatabase struct {
	*gorm.DB
}

func (db *GormDatabase) Where(query interface{}, args ...interface{}) Database {
	return &GormDatabase{db.DB.Where(query, args...)}
}

func (db *GormDatabase) First(dest interface{}, conds ...interface{}) Database {
	return &GormDatabase{db.DB.First(dest, conds...)}
}

func (db *GormDatabase) Error() error {
	return db.DB.Error
}

func NewDatabase() *gorm.DB {
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
			log.Printf("Attempt %d: Failed to initialize database. Retrying...", i)
			time.Sleep(3 * time.Second)
		}
	}

	migration.Migrate(database)

	return database
}
