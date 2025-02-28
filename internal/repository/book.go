package repository

import (
	"BookStore/internal/models"
)

type BookRepository interface {
	GetBooks() []models.Book
	GetBookByID(id int) (models.Book, error)
	GetBooksGroupedByCategoryRandom() (map[string][]models.Book, error)
}
