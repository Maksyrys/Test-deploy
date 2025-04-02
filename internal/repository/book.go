package repository

import (
	"BookStore/internal/models"
)

type BookRepository interface {
	GetBooks() []models.Book
	GetBookByID(id int) (models.Book, error)
	GetBooksGroupedByCategoryRandom() (map[string][]models.Book, error)
	GetBooksByGroupedByAuthorRandom() (map[string][]models.Book, error)
	SearchBooks(query string) ([]models.Book, error)
	GetRandomBooks(limit int) ([]models.Book, error)
	GetBooksByCategoryID(categoryID int) ([]models.Book, error)
	GetAllCategories() ([]models.Category, error)
	InsertBook(book models.Book) error
	DeleteBook(id int) error
}
