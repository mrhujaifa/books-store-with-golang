package app

import (
	"book-shop/internal/config"
	"book-shop/internal/health"
	"book-shop/internal/modules/auth"
	"book-shop/internal/modules/book"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func RegisterRoutes(cfg *config.Config, mux *http.ServeMux, db *gorm.DB) {
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Book Shop API is running"))
	})

	mux.HandleFunc("/health", health.HealthHandler) //* Health check endpoint

	bookModule := book.NewModule(db) //* Create a new instance of the book module with the database connection

	mux.HandleFunc("/api/v1/books", bookModule.Handler.CreateBook)
	mux.HandleFunc("/api/v1/all-books", bookModule.Handler.GetAllBooks)
	mux.HandleFunc("/api/v1/books/{id}", bookModule.Handler.GetBookById)
	mux.HandleFunc("/api/v1/books/delete/{id}", bookModule.Handler.DeleteBook)

	authModule, err := auth.NewModule(cfg, db)

	if err != nil {
		fmt.Println("hello world")
		panic("Failed to initialize Auth Module: " + err.Error())
	}

	mux.HandleFunc("/api/v1/auth/login", authModule.Handler.Login)
	mux.HandleFunc("/api/v1/auth/signup", authModule.Handler.Signup)
	mux.HandleFunc("/api/v1/auth/callback", authModule.Handler.Callback)
	mux.HandleFunc("/api/v1/auth/me", authModule.Handler.Me)
	mux.HandleFunc("/api/v1/auth/logout", authModule.Handler.Logout)

}
