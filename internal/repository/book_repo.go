package repository

import (
	"database/sql"
	"log"

	"github.com/seagilbert002/LittleLibrary/internal/models"
)

type BookRepository struct {
	DB *sql.DB
}

func NewSQLBookRepo(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

// FUNCTION: for returning all books in the database
func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	// SQL Query logic
	rows, err := r.DB.Query("SELECT id, title, author, publish_date, location FROM books")
	if err != nil {
		log.Printf("Repository Error: Failed to query books: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Load the books into a list we can display
    var books []models.Book
    for rows.Next() {
		var book models.Book

		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublishDate, &book.Location)
		if err != nil {
			log.Printf("Repository Error: Failed to scan book row: %v", err)
			return nil, err
		}
		books = append(books, book)
    }
	
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// FUNCTION: for returning a single book
func (r *BookRepository) GetBookById(id int) (*models.Book, error) {
	var book models.Book
	// SQL Query
	row := r.DB.QueryRow("SELECT title, author, first_name, last_name, genre, series, description, publish_date, publisher, ean_isbn, upc_isbn, pages, ddc, cover_style, sprayed_edges, special_ed, first_ed, signed, location FROM books WHERE id = ?", id)
	err := row.Scan(&book.Title, &book.Author, &book.AuthorFirst, &book.AuthorLast, &book.Genre, &book.Series, &book.Description, &book.PublishDate, &book.Publisher, &book.EanIsbn, &book.UpcIsbn, &book.Pages, &book.Ddc, &book.CoverStyle, &book.SprayedEdges, &book.SpecialEd, &book.FirstEd, &book.Signed, &book.Location)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

// FUNCTION: adding a book to the repo
func (r *BookRepository) AddBook(book models.Book) (error) {
	// Preps the statement for execution
	stmt, err := r.DB.Prepare("INSERT INTO books (title, author, first_name, last_name, genre, series, description, publish_date, publisher, ean_isbn, upc_isbn, pages, ddc, cover_style, sprayed_edges, special_ed, first_ed, signed, location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
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
		log.Printf("Book Insertion Failed: %v", err)
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
