package author

import (
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Author{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

func (b *BookRepository) FindAll() []Author {
	var books []Author
	b.db.Find(&books)

	return books
}

func (b *BookRepository) InsertSampleData(books []Author) {

	for _, book := range books {
		b.db.Where(Author{Name: book.Name}).
			Attrs(Author{Name: book.Name, ID: book.ID}).
			FirstOrCreate(&book)
	}
}
