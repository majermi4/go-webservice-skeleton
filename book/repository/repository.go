package repository

import (
	"MyWebService/book/data"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (br *BookRepository) Create(book *data.Book) error {
	return br.DB.Create(book).Error
}

func (br *BookRepository) Update(book *data.Book) error {
	return br.DB.Save(book).Error
}

func (br *BookRepository) Get(id string) *data.Book {
	var book *data.Book
	r := br.DB.Find(&book, "id = ?", id)

	if r.RowsAffected == 0 {
		return nil
	}

	return book
}

func (br *BookRepository) Delete(id string) error {
	return br.DB.Delete(&data.Book{}, "id = ?", id).Error
}

func (br *BookRepository) Find() ([]data.Book, error) {
	var books []data.Book
	result := br.DB.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}
