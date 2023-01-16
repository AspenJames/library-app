package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/aspenjames/library-app/server/models"
)

// Pointer to database connection.
var DB *gorm.DB

// Initialize database connection & automigrate models.
func Init() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	migrations := []interface{}{}
	for _, model := range models.Models {
		migrations = append(migrations, &model)
	}
	db.AutoMigrate(migrations...)
	DB = db
	return nil
}
