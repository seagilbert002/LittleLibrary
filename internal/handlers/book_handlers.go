package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"strconv"

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

// Handles deleting a book
func (h *BookHandler) RemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling request to: %s from %s", r.URL.Path, r.RemoteAddr)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	bookIDString := strings.TrimPrefix(r.URL.Path, "/remove_book/")
	bookId, err := strconv.Atoi(bookIDString)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = h.Catalog.RemoveBook(bookId)
	if err != nil {
        log.Printf("Service error deleting book: %v", err)
        http.Error(w, "Failed to delete book", http.StatusInternalServerError)
        return
    }

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}


// Handles displaying a single book
func (h *BookHandler) BookDisplayHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling request to: %s from: %s", r.URL.Path, r.RemoteAddr)

	// Get the requested book id
	bookIdString := strings.TrimPrefix(r.URL.Path, "/display_book/")
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Call the corresponding service
	book, err := h.Catalog.GetBookById(bookId)
	if err != nil {
		http.Error(w, "Failed to load book", http.StatusInternalServerError)
		return
	}

	// Render the template
	tmpl, err := template.ParseFiles("web/templates/pages/display_book.html")
	if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, book)
}

// Handles the Add Book Form and Posting the book
func (h *BookHandler) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling request to: %s from: %s", r.URL.Path, r.RemoteAddr)
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("web/templates/pages/add_book.html")
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Error rendering form", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error parsing bookform: %v ", err)
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		// Call the service that validates the book
		err = h.Catalog.AddBook(r.Form)

		if err != nil {
			http.Error(w, "Failed to add book", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handles the updating of a book and posting the book
func (h *BookHandler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling request to: %s from: %s", r.URL.Path, r.RemoteAddr)
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("web/templates/pages/update_book.html")
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Error rendering form", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error parsing bookform: %v ", err)
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		// Call the service that validates the book
		err = h.Catalog.UpdateBook(r.Form)

		if err != nil {
			http.Error(w, "Failed to update book", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
