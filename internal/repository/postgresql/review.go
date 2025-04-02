package postgresql

import (
	"BookStore/internal/models"
	"database/sql"
	"fmt"
)

type Review struct {
	DB *sql.DB
}

func NewReview(db *sql.DB) *Review {
	return &Review{DB: db}
}
func (r *Review) CreateReview(review models.Review) error {
	_, err := r.DB.Exec(`
        INSERT INTO reviews (user_id, book_id, rating, comment)
        VALUES ($1, $2, $3, $4)`,
		review.UserID, review.BookID, review.Rating, review.Comment)
	return err
}

func (r *Review) GetReviewsByBookID(bookID int) ([]models.Review, error) {
	rows, err := r.DB.Query(`
		SELECT r.review_id, r.user_id, u.username, r.book_id, r.rating, r.comment, r.created_at
		FROM reviews r
		JOIN users u ON r.user_id = u.user_id
		WHERE r.book_id = $1
		ORDER BY r.created_at DESC
	`, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(
			&review.ReviewID,
			&review.UserID,
			&review.Username,
			&review.BookID,
			&review.Rating,
			&review.Comment,
			&review.Created,
		); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, rows.Err()
}

func (r *Review) GetUserReviewsCount(userID int) (int, error) {
	var count int
	err := r.DB.QueryRow(`SELECT COUNT(*) FROM reviews WHERE user_id = $1`, userID).Scan(&count)
	return count, err
}

func (r *Review) GetUserReviews(userID int) ([]models.Review, error) {
	var reviews []models.Review

	rows, err := r.DB.Query(`
		SELECT r.review_id, r.user_id, u.username, r.book_id, b.title, r.rating, r.comment, r.created_at
		FROM reviews r
		JOIN users u ON r.user_id = u.user_id
		JOIN books b ON r.book_id = b.book_id
		WHERE r.user_id = $1
		ORDER BY r.created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var review models.Review
		err := rows.Scan(
			&review.ReviewID,
			&review.UserID,
			&review.Username,
			&review.BookID,
			&review.BookTitle,
			&review.Rating,
			&review.Comment, // <- добавьте поле comment
			&review.Created,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, rows.Err()
}

func (r *Review) UserHasReviewed(userID, bookID int) (bool, error) {
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM reviews WHERE user_id=$1 AND book_id=$2)", userID, bookID).Scan(&exists)
	return exists, err
}

func (r *Review) DeleteReview(reviewID int) error {
	query := "DELETE FROM reviews WHERE review_id = $1"
	_, err := r.DB.Exec(query, reviewID)
	return err
}

func (r *Review) GetAllReviews() ([]models.Review, error) {
	query := `
        SELECT review_id, user_id, book_id, rating, comment, created_at
        FROM reviews
        ORDER BY created_at DESC
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(
			&review.ReviewID,
			&review.UserID,
			&review.BookID,
			&review.Rating,
			&review.Comment,
			&review.Created,
		); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
		}
		reviews = append(reviews, review)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return reviews, nil
}
