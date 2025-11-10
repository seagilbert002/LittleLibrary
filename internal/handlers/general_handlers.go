package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// No dependencies yet but will add if needed
type GeneralHandler struct {}

// Constructor for the General Handler
func NewGeneralHandler() *GeneralHandler {
	return &GeneralHandler{}
}

// IndexHandler for the homepage
func (h *GeneralHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling request to: %s from %s", r.URL.Path, r.RemoteAddr)

	// Renders the index template
	tmpl, err := template.ParseFiles("web/templates/pages/index.html")
	if err != nil {
		log.Printf("Error rendering index template: %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
