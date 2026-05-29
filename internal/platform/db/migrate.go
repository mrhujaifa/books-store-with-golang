package db

import (
	"book-shop/internal/modules/auth"
	"book-shop/internal/modules/book"

	"gorm.io/gorm"
)

func RunMigrate(database *gorm.DB) error {
	return database.AutoMigrate(
		&book.Book{},
		&auth.User{},
	)
}
