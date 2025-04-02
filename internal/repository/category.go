package repository

import "BookStore/internal/models"

type CategoryRepository interface {
	CreateCategory(category models.Category) error
}
