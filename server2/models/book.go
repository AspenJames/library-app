package models

import (
	"fmt"

	"github.com/aspenjames/library-app/server/database"
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Base
	Edition    string `json:"edition,omitempty"`
	ISBN       string `json:"isbn,omitempty"`
	ReadStatus bool   `json:"read_status,omitempty"`
	Title      string `gorm:"not null" json:"title"`

	Authors []*Author `gorm:"many2many:book_authors" json:"authors,omitempty"`
}

func AllBooks() []Book {
	books := []Book{}
	database.DB.Preload("Authors").Find(&books)
	return books
}

func CreateBook(book *Book) error {
	return database.DB.Create(book).Error
}

func FindBook(id string) (Book, error) {
	book := Book{}
	dbID := IDFromString(id)
	database.DB.Preload("Authors").First(&book, dbID)
	if book.ID == UUIDNotFound {
		return book, fiber.NewError(404, fmt.Sprintf("book with ID %s not found", id))
	}
	return book, nil
}
