package book

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(book *Book) error
	GetAllBooks() ([]*Book, error)
	GetBookById(id string) (*Book, error)
	DeleteBook(id string) (*Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(book *Book) error {
	return r.db.Create(book).Error
}

func (r *repository) GetAllBooks() ([]*Book, error) {
	var books []*Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) GetBookById(id string) (*Book, error) {
	var book Book

	err := r.db.First(&book, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &book, nil
}

func (r *repository) DeleteBook(id string) (*Book, error) {
	var book Book

	err := r.db.Delete(&book, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return &book, nil
}
