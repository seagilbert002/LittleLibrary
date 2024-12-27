package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
    "strconv"
)

type Server struct {
    DB *sql.DB
}

type Book struct {
    Title           string
    Author          string
    AuthorFirst     string
    AuthorLast      string
    Genre           string
    Series          string
    Description     string
    PublishDate     string
    Publisher       string
    EanIsbn         string
    UpcIsbn         string
    Pages           int
    Ddc             string
    CoverStyle      string
    SprayedEdges    bool
    SpecialEd       bool
    FirstEd         bool
    Signed          bool
    Location        string
}

// A handler for the homepage
func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, nil)
}

// Handles the books page for displaying the books in the database
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
            "id":           fmt.Sprint(id),
            "title":        title,
            "author":       author,
            "publishDate":  publishDate,
            "location":     location,
        })
    }

    tmpl, err := template.ParseFiles("templates/books.html")
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, books)
}

// A handler for adding a book to the database
func (s *Server) AddBookHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        tmpl, err := template.ParseFiles("templates/add_book.html")
        if err != nil {
            http.Error(w, "Error rendering template", http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
    } else if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error parsing book form", http.StatusBadRequest)
            return
        }
        title :=        r.FormValue("title")
        author :=       r.FormValue("author")
        authorFirst :=  r.FormValue("first_name")
        authorLast :=   r.FormValue("last_name")
        genre :=        r.FormValue("genre")
        series :=       r.FormValue("series")
        description :=  r.FormValue("description")
        publishDate :=  r.FormValue("publish_date")
        publisher :=    r.FormValue("publisher")
        eanIsbn :=      r.FormValue("ean_isbn")
        upcIsbn :=      r.FormValue("upc_isbn")
        pagesStr :=     r.FormValue("pages")
        ddc :=          r.FormValue("ddc")
        coverStyle :=   r.FormValue("cover_style")
        sprayedEdges := r.FormValue("sprayed_edges") == "on"
        specialEd :=    r.FormValue("special_ed") == "on"
        firstEd :=      r.FormValue("first_ed") == "on"
        signed :=       r.FormValue("signed") == "on"
        location :=     r.FormValue("location")

        // Converting the pages to int
        pages, err := strconv.Atoi(pagesStr)
        if err != nil {
            http.Error(w, "Invalid pages value", http.StatusBadRequest)
            return
        }

        book := Book{
            Title:        title,
            Author:       author,
            AuthorFirst:  authorFirst,
            AuthorLast:   authorLast,
            Genre:        genre,
            Series:       series,
            Description:  description,
            PublishDate:  publishDate,
            Publisher:    publisher,
            EanIsbn:      eanIsbn,
            UpcIsbn:      upcIsbn,
            Pages:        pages,
            Ddc:          ddc,
            CoverStyle:   coverStyle,
            SprayedEdges: sprayedEdges,
            SpecialEd:    specialEd,
            FirstEd:      firstEd,
            Signed:       signed,
            Location:     location,
        }

        err = insertBook(s.DB, book)
        if err != nil {
            http.Error(w, "Error inserting book", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/books", http.StatusSeeOther)
    }
}

func insertBook(db *sql.DB, book Book) error {
    stmt, err := db.Prepare("INSERT INTO books (title, author, first_name, last_name, genre, series, description, publish_date, publisher, ean_isbn, upc_isbn, ddc, cover_style, sprayed_edges, special_ed, first_ed, signed, location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(book.Title, book.Author, book.AuthorFirst, book.AuthorLast, book.Genre, book.Series, book.Description, book.PublishDate, book.Publisher, book.EanIsbn, book.UpcIsbn, book.CoverStyle, book.SprayedEdges, book.SpecialEd, book.FirstEd, book.Signed, book.Location)
    return err
}
