package models

type Author struct {
	Base
	Name string `gorm:"not null" json:"name"`

	Books []*Book `gorm:"many2many:book_authors;" json:"-"`
}
