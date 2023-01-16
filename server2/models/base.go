package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Define core model fields.
type Base struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime:milli" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// UUID returned when an entity is not found.
var UUIDNotFound = uuid.MustParse("00000000-0000-0000-0000-000000000000")

// Return a UUID from a string.
func IDFromString(id string) uuid.UUID {
	return uuid.MustParse(id)
}
