package services

import (
	"BookStore/internal/models"
	"BookStore/internal/repository"
)

type BookService interface {
	GetIndexData(limit int) ([]models.Book, []models.Category, error)
	GetBookDetails(bookID int, user *models.User) (*BookDetails, error)
	SearchBooks(query string) ([]models.Book, error)
	AddBook(book models.Book) error
	DeleteBook(id int) error
	GetAllBooks() []models.Book
	CreateCategory(category models.Category) error
}

type bookService struct {
	rep *repository.Repository
}

func NewBookService(rep *repository.Repository) BookService {
	return &bookService{rep: rep}
}

type BookDetails struct {
	BooksByAuthor    map[string][]models.Book
	AuthorBooks      []models.Book
	Book             models.Book
	InCart           bool
	IsFavorite       bool
	BookReviews      []models.Review
	UserReviewExists bool
}

func (s *bookService) GetIndexData(limit int) ([]models.Book, []models.Category, error) {
	randomBooks, err := s.rep.Book.GetRandomBooks(limit)
	if err != nil {
		return nil, nil, err
	}
	categories, err := s.rep.Book.GetAllCategories()
	if err != nil {
		return nil, nil, err
	}
	return randomBooks, categories, nil
}

func (s *bookService) GetBookDetails(bookID int, user *models.User) (*BookDetails, error) {
	booksByAuthor, err := s.rep.Book.GetBooksByGroupedByAuthorRandom()
	if err != nil {
		return nil, err
	}

	book, err := s.rep.Book.GetBookByID(bookID)
	if err != nil {
		return nil, err
	}

	inCart := false
	isFavorite := false
	userReviewExists := false

	if user != nil {
		cartItems, err := s.rep.Cart.GetCartItems(user.UserId)
		if err == nil {
			for _, item := range cartItems {
				if item.BookID == book.ID {
					inCart = true
					break
				}
			}
		}
		isFavorite, err = s.rep.Favorite.IsFavorite(user.UserId, book.ID)
		if err != nil {

		}
		userReviewExists, _ = s.rep.Review.UserHasReviewed(user.UserId, book.ID)
	}

	authorBooks := booksByAuthor[book.Author]
	if len(authorBooks) > 6 {
		authorBooks = authorBooks[:6]
	}

	reviews, err := s.rep.Review.GetReviewsByBookID(book.ID)
	if err != nil {
		return nil, err
	}

	details := &BookDetails{
		BooksByAuthor:    booksByAuthor,
		AuthorBooks:      authorBooks,
		Book:             book,
		InCart:           inCart,
		IsFavorite:       isFavorite,
		BookReviews:      reviews,
		UserReviewExists: userReviewExists,
	}

	return details, nil
}

func (s *bookService) SearchBooks(query string) ([]models.Book, error) {
	return s.rep.Book.SearchBooks(query)
}

func (s *bookService) AddBook(book models.Book) error {
	return s.rep.Book.InsertBook(book)
}

func (s *bookService) DeleteBook(id int) error {
	return s.rep.Book.DeleteBook(id)
}

func (s *bookService) GetAllBooks() []models.Book {
	return s.rep.Book.GetBooks()
}

func (s *bookService) CreateCategory(category models.Category) error {
	return s.rep.Category.CreateCategory(category)
}
