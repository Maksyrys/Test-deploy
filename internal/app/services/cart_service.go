package services

import (
	"BookStore/internal/models"
	"BookStore/internal/repository"
)

type CartItemDetail struct {
	Book     models.Book
	Quantity int
	Total    float64
}

type CartDetails struct {
	Items         []CartItemDetail
	GrandTotal    float64
	TotalQuantity int
}

type CartService interface {
	GetCartDetails(userID int) (*CartDetails, error)
	AddItem(userID, bookID, quantity int) error
	RemoveOneItem(userID, bookID int) error
	RemoveAllItems(userID, bookID int) error
	GetCartCount(userID int) (int, error)
}

type cartService struct {
	rep *repository.Repository
}

func NewCartService(rep *repository.Repository) CartService {
	return &cartService{rep: rep}
}

func (s *cartService) GetCartDetails(userID int) (*CartDetails, error) {
	cartItems, err := s.rep.Cart.GetCartItems(userID)
	if err != nil {
		return nil, err
	}

	var details CartDetails

	for _, item := range cartItems {
		// Получаем данные книги по идентификатору.
		book, err := s.rep.Book.GetBookByID(item.BookID)
		if err != nil {
			// Если не удалось получить книгу, пропускаем элемент.
			continue
		}
		total := float64(item.Quantity) * book.Price
		details.GrandTotal += total
		details.TotalQuantity += item.Quantity

		details.Items = append(details.Items, CartItemDetail{
			Book:     book,
			Quantity: item.Quantity,
			Total:    total,
		})
	}

	return &details, nil
}

func (s *cartService) AddItem(userID, bookID, quantity int) error {
	return s.rep.Cart.AddItem(userID, bookID, quantity)
}

func (s *cartService) RemoveOneItem(userID, bookID int) error {
	return s.rep.Cart.RemoveOneItem(userID, bookID)
}

func (s *cartService) RemoveAllItems(userID, bookID int) error {
	return s.rep.Cart.RemoveItems(userID, bookID)
}

func (s *cartService) GetCartCount(userID int) (int, error) {
	cartItems, err := s.rep.Cart.GetCartItems(userID)
	if err != nil {
		return 0, err
	}
	totalCount := 0
	for _, item := range cartItems {
		totalCount += item.Quantity
	}
	return totalCount, nil
}
