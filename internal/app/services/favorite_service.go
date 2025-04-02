package services

import (
	"BookStore/internal/models"
	"BookStore/internal/repository"
)

type FavoriteService interface {
	// GetFavorites возвращает список избранных книг пользователя.
	GetFavorites(userID int) ([]models.Book, error)
	// AddFavorite добавляет книгу в избранное пользователя.
	AddFavorite(userID, bookID int) error
	// RemoveFavorite удаляет книгу из избранного пользователя.
	RemoveFavorite(userID, bookID int) error
	GetFavoritesCount(userID int) (int, error)
}

type favoriteService struct {
	rep *repository.Repository
}

func NewFavoriteService(rep *repository.Repository) FavoriteService {
	return &favoriteService{rep: rep}
}

func (s *favoriteService) GetFavorites(userID int) ([]models.Book, error) {
	return s.rep.Favorite.GetFavorites(userID)
}

func (s *favoriteService) AddFavorite(userID, bookID int) error {
	return s.rep.Favorite.AddFavorite(userID, bookID)
}

func (s *favoriteService) RemoveFavorite(userID, bookID int) error {
	return s.rep.Favorite.RemoveFavorite(userID, bookID)
}

func (s *favoriteService) GetFavoritesCount(userID int) (int, error) {
	favotiteItems, err := s.rep.Favorite.CountFavorites(userID)
	if err != nil {
		return 0, err
	}
	return favotiteItems, err
}
