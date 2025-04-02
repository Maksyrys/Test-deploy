package repository

import "BookStore/internal/models"

type FavoriteRepository interface {
	AddFavorite(userID, bookID int) error
	RemoveFavorite(userID, bookID int) error
	GetFavorites(userID int) ([]models.Book, error)
	IsFavorite(userID, bookID int) (bool, error)
	CountFavorites(userID int) (int, error)
}
