package postgresql

import (
	"BookStore/internal/models"
	"database/sql"
	"fmt"
	"strconv"
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
            b.detailed_description,   
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
		&book.DetailedDescription,
		&book.PublishDate,
		&book.ImageURL,
	)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *Book) GetBooksGroupedByCategoryRandom() (map[string][]models.Book, error) {

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

		result[categoryName] = append(result[categoryName], book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (b *Book) GetBooksByGroupedByAuthorRandom() (map[string][]models.Book, error) {
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
        ORDER BY RANDOM()
    `

	rows, err := b.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		var authorName string
		var categoryName string

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&authorName,
			&categoryName,
			&book.Price,
			&book.Description,
			&book.PublishDate,
			&book.ImageURL,
		)
		if err != nil {
			return nil, err
		}

		// Группировка по имени автора
		result[authorName] = append(result[authorName], book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (b *Book) SearchBooks(query string) ([]models.Book, error) {
	sqlQuery := `
    SELECT 
        b.book_id,
        b.title,
        a.name AS author_name,
        c.name AS category_name,
        b.price,
        b.description,
        b.detailed_description,
        b.publish_date,
        b.image_url
    FROM books b
    LEFT JOIN authors a ON b.author_id = a.author_id
    LEFT JOIN categories c ON b.category_id = c.category_id
    WHERE b.title ILIKE $1
       OR a.name ILIKE $1
    ORDER BY b.title ASC
    `

	rows, err := b.DB.Query(sqlQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Category,
			&book.Price,
			&book.Description,
			&book.DetailedDescription,
			&book.PublishDate,
			&book.ImageURL,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (b *Book) GetRandomBooks(limit int) ([]models.Book, error) {
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
		ORDER BY RANDOM()
		LIMIT $1
	`

	rows, err := b.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Category,
			&book.Price,
			&book.Description,
			&book.PublishDate,
			&book.ImageURL,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (b *Book) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	rows, err := b.DB.Query("SELECT category_id, name FROM categories ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, rows.Err()
}

func (b *Book) GetBooksByCategoryID(categoryID int) ([]models.Book, error) {
	query := `
        SELECT 
            b.book_id,
            b.title,
            a.name AS author_name,
            c.name AS category_name,
            b.price,
            b.description,
            b.detailed_description,
            b.publish_date,
            b.image_url
        FROM books b
        LEFT JOIN authors a ON b.author_id = a.author_id
        LEFT JOIN categories c ON b.category_id = c.category_id
        WHERE b.category_id = $1
        ORDER BY b.title
    `
	rows, err := b.DB.Query(query, categoryID)
	if err != nil {
		return nil, err
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
			&book.DetailedDescription,
			&book.PublishDate,
			&book.ImageURL,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, rows.Err()
}

func (b *Book) InsertBook(book models.Book) error {
	query := `
        INSERT INTO books 
            (title, author_id, category_id, price, description, detailed_description, publish_date, image_url)
        VALUES 
            ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	var authorID interface{}
	if book.Author != "" {
		if id, err := strconv.Atoi(book.Author); err == nil {
			authorID = id
		} else {
			return fmt.Errorf("некорректное значение для author_id: %v", err)
		}
	} else {
		authorID = nil
	}

	var categoryID interface{}
	if book.Category != "" {
		if id, err := strconv.Atoi(book.Category); err == nil {
			categoryID = id
		} else {
			return fmt.Errorf("некорректное значение для category_id: %v", err)
		}
	} else {
		categoryID = nil
	}

	_, err := b.DB.Exec(query,
		book.Title,
		authorID,
		categoryID,
		book.Price,
		book.Description,
		book.DetailedDescription,
		book.PublishDate,
		book.ImageURL,
	)
	return err
}

func (b *Book) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE book_id = $1"
	_, err := b.DB.Exec(query, id)
	return err
}
