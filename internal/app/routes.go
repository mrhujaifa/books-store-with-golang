package app

import (
	"book-shop/internal/health"
	"book-shop/internal/modules/book"
	"net/http"

	"gorm.io/gorm"
)

func RegisterRoutes(mux *http.ServeMux, db *gorm.DB) {
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Book Shop API is running"))
	})

	mux.HandleFunc("/health", health.HealthHandler) //* Health check endpoint

	bookModule := book.NewModule(db) //* Create a new instance of the book module with the database connection

	mux.HandleFunc("/api/v1/books", bookModule.Handler.CreateBook)
	mux.HandleFunc("/api/v1/all-books", bookModule.Handler.GetAllBooks)
	mux.HandleFunc("/api/v1/books/{id}", bookModule.Handler.GetBookById)
	mux.HandleFunc("/api/v1/books/delete/{id}", bookModule.Handler.DeleteBook)
}
