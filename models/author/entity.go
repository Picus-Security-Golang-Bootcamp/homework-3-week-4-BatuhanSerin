package author

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	//Books []book.BookSlice
	ID         string
	Name       string
	Page       string
	Stock      string
	Cost       string
	StockCode  string
	ISBN       string
	AuthorID   string
	AuthorName string
}
type AuthorSlice []Author
