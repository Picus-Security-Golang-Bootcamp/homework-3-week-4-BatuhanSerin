package author

import (
	"errors"
	"fmt"
	"strconv"

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

func (b *BookRepository) InsertData(books []Author) {

	for _, book := range books {
		b.db.Where(Author{Name: book.Name}).Attrs(Author{Name: book.Name, ID: book.ID}).FirstOrCreate(&book)
	}
}
func (b *BookRepository) GetByID(id int) (*Author, error) {
	var book Author

	result := b.db.First(&book, strconv.Itoa(id))
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	fmt.Println("The book with id :", id)

	return &book, nil
}
func (b *BookRepository) FindByName(name string) []Author {
	var books []Author
	b.db.Where("name LIKE ? ", "%"+name+"%").Find(&books)
	fmt.Println("The book with named as", name)

	return books
}
func (b *BookRepository) GetBooksWithAuthor(authorName string) []Author {
	var books []Author
	b.db.Where("author_name LIKE ? ", "%"+authorName+"%").Find(&books)
	fmt.Println("The book with Author name (", authorName, ")")
	return books

}

func (b *BookRepository) DeleteById(id int) error {
	result := b.db.Delete(&Author{}, strconv.Itoa(id))

	if result.Error != nil {
		return result.Error
	}

	return nil
}
