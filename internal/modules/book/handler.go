package book

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateBook(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Method not allowed",
		})
		return
	}

	var input CreateBookInput

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}
	book, err := h.service.CreateBook(input)

	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Failed to create book",
		})
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"message": "Book created successfully",
		"data":    book,
	})

}

// & Handler for getting books
func (h *Handler) GetAllBooks(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Method not allowed",
		})

		return
	}

	books, err := h.service.GetAllBooks()
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Failed to get books",
		})
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(map[string]interface{}{
		"message": "Books retrieved successfully",
		"data":    books,
	})

}

// & Handler for getting books by ID
func (h *Handler) GetBookById(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Method not allowed",
		})

	}

	bookId := req.PathValue("id")

	book, err := h.service.GetBookById(bookId)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]interface{}{
			"message": "Failed to get book by id",
			"success": false,
			"status":  http.StatusInternalServerError,
		})
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Get book by id fetched successfully",
		"success": true,
		"data":    book,
	})
}

// & Handler for delete book by ID
func (h *Handler) DeleteBook(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Method not allowed",
		})
		return
	}

	bookId := req.PathValue("id")

	book, err := h.service.DeleteBook(bookId)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(map[string]interface{}{
			"message": "Failed to delete book by id",
			"success": false,
			"status":  http.StatusInternalServerError,
		})
		return
	}

	if book == nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]interface{}{
			"status":  http.StatusNotFound,
			"message": "No book found! Please try again later.",
			"success": false,
		})
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Book deleted successfully",
		"success": true,
	})
}
