package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

)

type Server struct {
    DB *sql.DB
}

type Book struct {
    title           string
    author          string
    firstName       string
    lastName        string
    genre           string
    series          string
    description     string
    publishDate     string
    publisher       string
    ean_isbn        string
    upc_isbn        string
    pages           int32
    ddc             string
    coverStyle      string
    sprayedEdges    bool
    specialEd       bool
    firstEd         bool
    signed          bool
    location        bool
}

// A handler for the homepage
func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, nil)
}

// Handles the books page for inserting a new book or more
func (s *Server) BooksHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := s.DB.Query("SELECT id, title, author, publish_date, location FROM books")
    if err != nil {
        http.Error(w, "Error fetching books", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var books []map[string]string
    for rows.Next() {
        var id int
        var title, author, publishDate, location string
        rows.Scan(&id, &title, &author, &publishDate, &location)

        books = append(books, map[string]string {
            "id": fmt.Sprint(id),
            "title": title,
            "author": author,
            "publishDate": publishDate,
            "location": location,
        })
    }

    tmpl, err := template.ParseFiles("templates/books.html")
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, books)
}