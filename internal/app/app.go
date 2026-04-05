package app

import (
	"net/http"

	"book-shop/internal/config"

	"gorm.io/gorm"
)

type App struct {
	server *http.Server
}

func NewApp(cfg *config.Config, db *gorm.DB) *App {
	mux := http.NewServeMux()

	RegisterRoutes(mux, db)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	return &App{
		server: server,
	}
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
