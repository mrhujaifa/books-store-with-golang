package book

type CreateBookInput struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
