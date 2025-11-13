package services

import "github.com/seagilbert002/LittleLibrary/internal/models"

// Repository interface for calling to the database
type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	// TODO: AddBook
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
