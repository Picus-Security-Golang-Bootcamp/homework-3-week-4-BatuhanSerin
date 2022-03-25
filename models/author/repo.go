package author

import (
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Author{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}
