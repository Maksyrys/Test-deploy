package repository

import "BookStore/internal/models"

type ReviewRepository interface {
	CreateReview(review models.Review) error
	GetReviewsByBookID(bookID int) ([]models.Review, error)
	GetUserReviewsCount(userID int) (int, error)
	GetUserReviews(userID int) ([]models.Review, error)
	UserHasReviewed(userID, bookID int) (bool, error)
	DeleteReview(reviewID int) error
	GetAllReviews() ([]models.Review, error)
}
