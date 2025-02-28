package repository

import (
	"BookStore/internal/repository/postgresql"
	"database/sql"
)

type Repository struct {
	Book BookRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Book: postgresql.NewBook(db),
	}
}
