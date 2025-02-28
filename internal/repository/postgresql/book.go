package postgresql

import (
	"BookStore/internal/models"
	"database/sql"
)

type Book struct {
	DB *sql.DB
}

func NewBook(db *sql.DB) *Book {
	return &Book{DB: db}
}

func (b *Book) GetBooks() []models.Book {
	query := `
        SELECT 
            b.book_id, 
            b.title, 
            a.name AS author_name,
            c.name AS category_name,
            b.price, 
            b.description, 
            b.publish_date,
            b.image_url
        FROM books b
        LEFT JOIN authors a ON b.author_id = a.author_id
        LEFT JOIN categories c ON b.category_id = c.category_id
    `
	rows, err := b.DB.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Category,
			&book.Price,
			&book.Description,
			&book.PublishDate,
			&book.ImageURL,
		); err != nil {
			return nil
		}
		books = append(books, book)
	}
	return books
}

func (b *Book) GetBookByID(id int) (models.Book, error) {
	var book models.Book
	query := `
        SELECT 
            b.book_id, 
            b.title, 
            a.name AS author_name,
            c.name AS category_name,
            b.price, 
            b.description,
            b.detailed_description,   -- !!! Добавляем
            b.publish_date,
            b.image_url
        FROM books b
        LEFT JOIN authors a ON b.author_id = a.author_id
        LEFT JOIN categories c ON b.category_id = c.category_id
        WHERE b.book_id = $1
        LIMIT 1
    `

	row := b.DB.QueryRow(query, id)
	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Category,
		&book.Price,
		&book.Description,
		&book.DetailedDescription, // считываем
		&book.PublishDate,
		&book.ImageURL,
	)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *Book) GetBooksGroupedByCategoryRandom() (map[string][]models.Book, error) {
	// В результат будем складывать книги, сгруппированные по названию категории
	result := make(map[string][]models.Book)

	query := `
        SELECT 
            b.book_id,
            b.title,
            a.name AS author_name,
            c.name AS category_name,
            b.price,
            b.description,
            b.publish_date,
            b.image_url
        FROM books b
        LEFT JOIN authors a ON b.author_id = a.author_id
        LEFT JOIN categories c ON b.category_id = c.category_id
        -- ORDER BY RANDOM() - для PostgreSQL вернёт строки в случайном порядке
        ORDER BY RANDOM()
    `

	rows, err := b.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		var categoryName string

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&categoryName,
			&book.Price,
			&book.Description,
			&book.PublishDate,
			&book.ImageURL,
		)
		if err != nil {
			return nil, err
		}

		// Добавляем книгу в срез в map по ключу = categoryName
		result[categoryName] = append(result[categoryName], book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
