package author

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        string
	Name      string
	Page      string
	Stock     string
	Cost      string
	StockCode string
	ISBN      string
}
type BookSlice []Book
