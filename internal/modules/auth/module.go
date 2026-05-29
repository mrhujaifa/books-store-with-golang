package auth

import (
	config "book-shop/internal/config"

	"gorm.io/gorm"
)

type Module struct {
	Handler *Handler
}

func NewModule(cfg *config.Config, db *gorm.DB) (*Module, error) {
	repo := NewRepository(db)
	service := NewService(repo)

	oidcManager, err := NewOIDCManager(
		cfg.Auth0Domain,
		cfg.Auth0ClientID,
		cfg.Auth0ClientSecret,
		cfg.Auth0CallbackURL,
	)
	if err != nil {
		return nil, err
	}

	store := NewSessionStore(cfg.SessionSecret)

	handler := NewHandler(service, oidcManager, store, cfg.AppBaseURL)

	return &Module{
		Handler: handler,
	}, nil
}
