package book

import "gorm.io/gorm"

type Module struct {
	Handler *Handler
}

func NewModule(db *gorm.DB) *Module {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	return &Module{
		Handler: handler,
	}
}
