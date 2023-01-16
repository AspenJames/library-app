package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Pointer to database connection.
var DB *gorm.DB

// Run AutoMigrate against DB.
func AutoMigrate(migrations []interface{}) error {
	return DB.AutoMigrate(migrations...)
}

// Initialize database connection & automigrate models.
func Init() error {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}
