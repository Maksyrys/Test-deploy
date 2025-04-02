package repository

import (
	"BookStore/internal/repository/postgresql"
	"database/sql"
)

type Repository struct {
	Favorite FavoriteRepository
	User     UserRepository
	Book     BookRepository
	Cart     CartRepository
	Review   ReviewRepository
	Category CategoryRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Favorite: postgresql.NewFavorite(db),
		Book:     postgresql.NewBook(db),
		User:     postgresql.NewUser(db),
		Cart:     postgresql.NewCart(db),
		Review:   postgresql.NewReview(db),
		Category: postgresql.NewCategory(db),
	}
}
