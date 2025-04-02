package repository

import "BookStore/internal/models"

type CartRepository interface {
	AddItem(userID, bookID, quantity int) error
	RemoveOneItem(userID, bookID int) error
	GetCartItems(userID int) ([]models.CartItem, error)
	ClearCart(userID int) error
	RemoveItems(userID, bookID int) error
	CountItems(userID int) (int, error)
}
