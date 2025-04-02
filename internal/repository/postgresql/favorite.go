package postgresql

import (
	"BookStore/internal/models"
	"database/sql"
)

type Favorite struct {
	DB *sql.DB
}

func NewFavorite(db *sql.DB) *Favorite {
	return &Favorite{DB: db}
}

func (f *Favorite) AddFavorite(userID, bookID int) error {
	// Проверяем, есть ли уже такая запись
	var exists bool
	err := f.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id=$1 AND book_id=$2)", userID, bookID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return nil // уже добавлено
	}
	_, err = f.DB.Exec("INSERT INTO favorites (user_id, book_id, added_at) VALUES ($1, $2, NOW())", userID, bookID)
	return err
}

func (f *Favorite) RemoveFavorite(userID, bookID int) error {
	_, err := f.DB.Exec("DELETE FROM favorites WHERE user_id=$1 AND book_id=$2", userID, bookID)
	return err
}

func (f *Favorite) GetFavorites(userID int) ([]models.Book, error) {
	query := `
    SELECT b.book_id, b.title, a.name as author_name, c.name as category_name, 
           b.price, b.description, b.publish_date, b.image_url
    FROM favorites f
    JOIN books b ON f.book_id = b.book_id
    LEFT JOIN authors a ON b.author_id = a.author_id
    LEFT JOIN categories c ON b.category_id = c.category_id
    WHERE f.user_id = $1
    ORDER BY f.added_at DESC`
	rows, err := f.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Price, &book.Description, &book.PublishDate, &book.ImageURL); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, rows.Err()
}

func (f *Favorite) IsFavorite(userID, bookID int) (bool, error) {
	var exists bool
	err := f.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id=$1 AND book_id=$2)", userID, bookID).Scan(&exists)
	return exists, err
}

func (f *Favorite) CountFavorites(userID int) (int, error) {
	var count int
	err := f.DB.QueryRow("SELECT COUNT(*) FROM favorites WHERE user_id=$1", userID).Scan(&count)
	return count, err
}
