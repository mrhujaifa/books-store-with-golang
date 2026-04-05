package book

import "time"

type Book struct {
	ID          string    `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
