package models

type Book struct {
	Base
	Edition    string
	ISBN       string
	ReadStatus bool
	Title      string

	Authors []*Author `gorm:"many2many:book_authors"`
}
