package repository

import (
	"database/sql"
	"log"
	"github.com/seagilbert002/LittleLibrary/internal/models"
)

type BookRepository struct {
	db *sql.DB
}

func (r *BookRepository) DisplayAllBooks() ([]models.Book, error) {
	// SQL Query logic
	rows, err := r.db.Query("SELECT id, title, author, publish_date, location FROM books")
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
