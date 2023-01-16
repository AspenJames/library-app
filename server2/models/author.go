package models

type Author struct {
	Base
	Name string

	Books []*Book `gorm:"many2many:book_authors;"`
}
