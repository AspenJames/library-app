package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Define core model fields
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
