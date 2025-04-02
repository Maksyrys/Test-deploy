package services

import (
	"BookStore/internal/models"
	"BookStore/internal/repository"
	"fmt"
	"time"
)

type ReviewService interface {
	CreateReview(review models.Review) error
	DeleteReview(reviewID int) error
	GetAllReviews() ([]models.Review, error)
}

type reviewService struct {
	rep *repository.Repository
}

func NewReviewService(rep *repository.Repository) ReviewService {
	return &reviewService{rep: rep}
}

func (s *reviewService) CreateReview(review models.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return fmt.Errorf("некорректный рейтинг: %d", review.Rating)
	}

	if review.Created.IsZero() {
		review.Created = time.Now()
	}
	return s.rep.Review.CreateReview(review)
}

func (s *reviewService) DeleteReview(reviewID int) error {
	return s.rep.Review.DeleteReview(reviewID)
}

func (s *reviewService) GetAllReviews() ([]models.Review, error) {
	return s.rep.Review.GetAllReviews()
}
