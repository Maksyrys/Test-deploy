package repository

import "BookStore/internal/models"

type UserRepository interface {
	CreateUser(user models.User) (int, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	UpdateUserProfile(user models.User) error
}
