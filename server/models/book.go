package models

import (
	"fmt"

	"github.com/aspenjames/library-app/server/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	Base
	Edition    string `json:"edition,omitempty"`
	ISBN       string `json:"isbn,omitempty"`
	ReadStatus bool   `json:"read_status,omitempty"`
	Title      string `gorm:"not null" json:"title"`

	Authors []*Author `gorm:"many2many:book_authors" json:"authors,omitempty"`
}

func AllBooks() ([]Book, error) {
	books := []Book{}
	res := database.DB.Preload("Authors").Find(&books)
	return books, res.Error
}

func CreateBook(book *Book) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// Ensure we FindOrCreate Authors by Name
		if len(book.Authors) > 0 {
			for _, author := range book.Authors {
				if err := FindOrCreateAuthor(author, tx); err != nil {
					return err
				}
			}
		}
		return tx.Create(book).Error
	})
}

func DeleteBook(id string) error {
	return database.DB.Delete(&Book{}, IDFromString(id)).Error
}

func FindBook(id string) (Book, error) {
	book := Book{}
	res := database.DB.Preload("Authors").First(&book, IDFromString(id))
	if res.Error != nil {
		return book, res.Error
	}
	if book.ID == UUIDNotFound {
		return book, fiber.NewError(404, fmt.Sprintf("book with ID '%s' not found", id))
	}
	return book, nil
}

func UpdateBook(id string, book *Book) (Book, error) {
	found := Book{}
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// Ensure Book exists first
		res := tx.Find(&found, IDFromString(id))
		if res.Error != nil {
			return res.Error
		}
		if found.ID == UUIDNotFound {
			return fiber.NewError(404, fmt.Sprintf("book with id '%s' not found", book.ID))
		}
		// Ensure we FindOrCreate Authors by Name
		if len(book.Authors) > 0 {
			for _, author := range book.Authors {
				if err := FindOrCreateAuthor(author, tx); err != nil {
					return err
				}
			}
			if err := tx.Model(&found).Association("Authors").Replace(book.Authors); err != nil {
				return err
			}
		}
		return tx.Model(&found).Updates(book).Error
	})
	return found, err
}
