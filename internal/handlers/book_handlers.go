package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/seagilbert002/LittleLibrary/internal/services"
)

// Book Handler will require a service
type BookHandler struct {
	Catalog *services.CatalogService // depends on the service layer
}

// NewBookHandler creates a new handler instance
func NewBookHandler(s *services.CatalogService) *BookHandler {
	return &BookHandler{Catalog: s}
}

// TODO: Separate db operations from BooksHandler
// Handles the books page for displaying the books in the database
func  (h *BookHandler) BooksHanlder (w http.ResponseWriter, r *http.Request) {
    log.Printf("Handling request to: %s from: %s", r.URL.Path, r.RemoteAddr)
	
	// Call the service
	books, err := h.Catalog.GetAllBooks()
	if err != nil {
		http.Error(w, "Failed to loade book catalog", http.StatusInternalServerError)
		return
	}

    tmpl, err := template.ParseFiles("web/templates/pages/books.html")
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, books)
}


