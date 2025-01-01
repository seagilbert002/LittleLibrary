package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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
    Pages           uint16 
    Ddc             string
    CoverStyle      string
    SprayedEdges    bool
    SpecialEd       bool
    FirstEd         bool
    Signed          bool
    Location        string
}

// A handler for the homepage
func IndexHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    log.Printf("Handling request to: %s from: %s", r.URL.Path, r.RemoteAddr)
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, nil)
}

// Handles the books page for displaying the books in the database
func  BooksHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    log.Printf("Handling request to: %s from: %s", r.URL.Path, r.RemoteAddr)
    rows, err := db.Query("SELECT id, title, author, publish_date, location FROM books")
    if err != nil {
        http.Error(w, "Error fetching books", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var books []map[string]string
    for rows.Next() {
        var id, title, author, publishDate, location string
        rows.Scan(&id, &title, &author, &publishDate, &location)

        books = append(books, map[string]string {
            "id":           id,
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
func AddBookHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    log.Printf("Handling request to: %s from: %s", r.URL.Path, r.RemoteAddr)
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
            log.Printf("Error parsing bookform: %v ", err)
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

        // Validation and Default values
        var errors []string

        if strings.TrimSpace(title) == "" {
            errors = append(errors, "Title is required")
        }
        if strings.TrimSpace(author) == "" {
            errors = append(errors, "Author is required") 
        }

        // Converting the pages to int
        pages, err := strconv.ParseUint(pagesStr, 10, 16)
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
            Pages:        uint16(pages),
            Ddc:          ddc,
            CoverStyle:   coverStyle,
            SprayedEdges: sprayedEdges,
            SpecialEd:    specialEd,
            FirstEd:      firstEd,
            Signed:       signed,
            Location:     location,
        }

        err = insertBook(db, book)
        if err != nil {
            http.Error(w, "Error inserting book", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/books", http.StatusSeeOther)
    }
}

func insertBook(db *sql.DB, book Book) error {
    stmt, err := db.Prepare("INSERT INTO books (title, author, first_name, last_name, genre, series, description, publish_date, publisher, ean_isbn, upc_isbn, pages, ddc, cover_style, sprayed_edges, special_ed, first_ed, signed, location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
        log.Printf("Error preparing statement: %v", err)
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(
        book.Title, 
        book.Author, 
        book.AuthorFirst, 
        book.AuthorLast, 
        book.Genre, 
        book.Series, 
        book.Description, 
        book.PublishDate, 
        book.Publisher, 
        book.EanIsbn, 
        book.UpcIsbn, 
        book.Pages, 
        book.Ddc,
        book.CoverStyle, 
        book.SprayedEdges, 
        book.SpecialEd, 
        book.FirstEd, 
        book.Signed, 
        book.Location,
    )
    if err != nil {
        log.Printf("Book Insert Failed: %v", err)
        return err
    }

    insertedId, err := result.LastInsertId()
    if err != nil {
        log.Printf("Error getting inserted ID: %v", err)
        return err
    }
    log.Printf("Book Inserted with ID: %d", insertedId)
    return nil
}

func BookDisplayHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    log.Printf("Handling request to: %s from: %s", r.URL.Path, r.RemoteAddr)
    bookIdStr := strings.TrimPrefix(r.URL.Path, "/display_book/")
    bookId, err := strconv.Atoi(bookIdStr)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    row := db.QueryRow("SELECT title, author, first_name, last_name, genre, series, description, publish_date, publisher, ean_isbn, upc_isbn, pages, ddc, cover_style, sprayed_edges, special_ed, first_ed, signed, location FROM books WHERE id = ?", bookId)
    var book Book
    err = row.Scan(&book.Title, &book.Author, &book.AuthorFirst, &book.AuthorLast, &book.Genre, &book.Series, &book.Description, &book.PublishDate, &book.Publisher, &book.EanIsbn, &book.UpcIsbn, &book.Pages, &book.Ddc, &book.CoverStyle, &book.SprayedEdges, &book.SpecialEd, &book.FirstEd, &book.Signed, &book.Location)
    if err == sql.ErrNoRows {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    } else if err != nil {
        log.Printf("Error querying database: %v", err)
        http.Error(w, "Database error", http.StatusBadRequest)
        return
    }

    tmpl, err := template.ParseFiles("templates/display_book.html") // Parse the detail template
    if err != nil {
        log.Printf("Error parsing template: %v", err)
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, book) // Execute the template with book data
    if err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Error executing template", http.StatusInternalServerError)
        return
    }
}
