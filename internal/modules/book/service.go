package book

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	CreateBook(input CreateBookInput) (*Book, error)
	GetAllBooks() ([]*Book, error)
	GetBookById(id string) (*Book, error)
	DeleteBook(id string) (*Book, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateBook(input CreateBookInput) (*Book, error) {
	if input.Title == "" {
		return nil, errors.New("title is required")
	}

	if input.Author == "" {
		return nil, errors.New("author is required")
	}

	if input.Price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}

	book := &Book{
		ID:          uuid.New().String(),
		Title:       input.Title,
		Author:      input.Author,
		Price:       int(input.Price),
		Description: input.Description,
	}

	err := s.repo.Create(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// & Getting books service method
func (s *service) GetAllBooks() ([]*Book, error) {
	books, err := s.repo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

// & Getting books by ID
func (s *service) GetBookById(id string) (*Book, error) {
	book, err := s.repo.GetBookById(id)

	if err != nil {
		return nil, err
	}

	return book, err
}

// & Delete book
func (s *service) DeleteBook(id string) (*Book, error) {
	book, err := s.repo.DeleteBook(id)

	if err != nil {
		return nil, err
	}

	return book, err
}
