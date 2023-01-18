package models

import (
	"fmt"

	"github.com/aspenjames/library-app/server/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Author struct {
	Base
	Name string `gorm:"not null;unique" json:"name"`

	Books []*Book `gorm:"many2many:book_authors;" json:"-"`
}

func AllAuthors() ([]Author, error) {
	authors := []Author{}
	res := database.DB.Find(&authors)
	return authors, res.Error
}

func FindAuthor(id string) (Author, error) {
	author := Author{}
	res := database.DB.First(&author, IDFromString(id))
	if res.Error != nil {
		return author, res.Error
	}
	if author.ID == UUIDNotFound {
		return author, fiber.NewError(404, fmt.Sprintf("author with ID '%s' not found", id))
	}
	return author, nil
}

// Find Author by Name, creating if not exists. Accepts a DB conn or trx.
func FindOrCreateAuthor(author *Author, db *gorm.DB) error {
	return db.Where(Author{Name: author.Name}).FirstOrCreate(author).Error
}
