package postgresql

import (
	"BookStore/internal/models"
	"database/sql"
	"fmt"
)

type Category struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db}
}

func (c *Category) CreateCategory(category models.Category) error {
	query := "INSERT INTO categories (name) VALUES $1"

	_, err := c.db.Exec(query, category.Name)
	if err != nil {
		return fmt.Errorf("Ошибка добавления новой категории: %s", err)
	}
	return nil
}
