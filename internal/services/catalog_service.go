package services

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/seagilbert002/LittleLibrary/internal/models"
)

// Repository interface for calling to the database
type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	AddBook(models.Book) (error)
	GetBookById(id int) (*models.Book, error)
}

// Catalog Service struct for business logic
type CatalogService struct {
	Repo BookRepository
}

// Wires together the dependencies
func NewCatalogService(repo BookRepository) *CatalogService {
	return &CatalogService{Repo: repo}
}

func (s *CatalogService) GetAllBooks() ([]models.Book, error) {
	return s.Repo.GetAllBooks()
}

func (s *CatalogService) GetBookById(id int) (*models.Book, error) {
	// Future authorization checks

	// Pull from the repository
	return s.Repo.GetBookById(id)
}

func (s *CatalogService) AddBook(bookData url.Values) error {
	// Validate the required fields
	if bookData.Get("title") == "" {
		return errors.New("book title is required")
	}
	if bookData.Get("pages") == "" {
		return errors.New("book page numbers is required")
	}
	if bookData.Get("location") == "" {
		return errors.New("book location is required")
	}

	// Transform data if needed here
	pagesStr := bookData.Get("pages")
	var pages uint16

	if pagesStr != "" {
		p, err := strconv.ParseUint(pagesStr, 10, 16)
		if err != nil {
			return errors.New("pages must be a valid number")
		}
		pages = uint16(p)
	}

	// Assembling the Model
	book := models.Book{
		Title: 			bookData.Get("title"),
		Author:			bookData.Get("author"),
		AuthorFirst:	bookData.Get("first_name"),
		AuthorLast:		bookData.Get("last_name"),
		Genre:			bookData.Get("genre"),
		Series:			bookData.Get("series"),
		Description:	bookData.Get("description"),
		PublishDate:	bookData.Get("publish_date"),
		Publisher:		bookData.Get("publisher"),
		EanIsbn:		bookData.Get("ean_isbn"),
		UpcIsbn:		bookData.Get("upc_isbn"),
		Pages:			pages,
		Ddc:			bookData.Get("ddc"),
		CoverStyle:		bookData.Get("cover_style"),
		SprayedEdges:	bookData.Get("sprayed_edges") == "on",
		SpecialEd:		bookData.Get("special_ed") == "on",
		FirstEd:		bookData.Get("first_ed") == "on",
		Signed:			bookData.Get("signed") == "on",
		Location:		bookData.Get("location"),
	}

	return s.Repo.AddBook(book)
}
