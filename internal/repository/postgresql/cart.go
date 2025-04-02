package postgresql

import (
	"BookStore/internal/models"
	"database/sql"
)

type Cart struct {
	DB *sql.DB
}

func NewCart(db *sql.DB) *Cart {
	return &Cart{DB: db}
}

func (c *Cart) AddItem(userID, bookID, quantity int) error {
	// Если элемент уже есть в корзине, можно обновить количество
	var currentQuantity int
	err := c.DB.QueryRow("SELECT quantity FROM cart_items WHERE user_id = $1 AND book_id = $2", userID, bookID).Scan(&currentQuantity)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		// Вставляем новую запись
		_, err = c.DB.Exec("INSERT INTO cart_items (user_id, book_id, quantity) VALUES ($1, $2, $3)", userID, bookID, quantity)
		return err
	}

	// Если запись уже существует – обновляем количество
	_, err = c.DB.Exec("UPDATE cart_items SET quantity = $1 WHERE user_id = $2 AND book_id = $3", currentQuantity+quantity, userID, bookID)
	return err
}

func (c *Cart) RemoveItems(userID, bookID int) error {
	_, err := c.DB.Exec("DELETE FROM cart_items WHERE user_id = $1 AND book_id = $2", userID, bookID)
	return err
}

func (c *Cart) RemoveOneItem(userID, bookID int) error {
	// Получаем текущее количество товара
	var currentQuantity int
	err := c.DB.QueryRow("SELECT quantity FROM cart_items WHERE user_id = $1 AND book_id = $2", userID, bookID).Scan(&currentQuantity)
	if err != nil {
		return err
	}

	if currentQuantity > 1 {
		// Если книг больше одной, уменьшаем количество на 1
		_, err = c.DB.Exec("UPDATE cart_items SET quantity = quantity - 1 WHERE user_id = $1 AND book_id = $2", userID, bookID)
		return err
	}
	// Если количество равно 1, удаляем запись
	_, err = c.DB.Exec("DELETE FROM cart_items WHERE user_id = $1 AND book_id = $2", userID, bookID)
	return err
}

func (c *Cart) GetCartItems(userID int) ([]models.CartItem, error) {
	rows, err := c.DB.Query("SELECT cart_item_id, user_id, book_id, quantity, added_at FROM cart_items WHERE user_id = $1 ORDER BY added_at ASC", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.CartItem
	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.CartItemID, &item.UserID, &item.BookID, &item.Quantity, &item.AddedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (c *Cart) ClearCart(userID int) error {
	_, err := c.DB.Exec("DELETE FROM cart_items WHERE user_id = $1", userID)
	return err
}

func (c *Cart) CountItems(userID int) (int, error) {
	var totalQuantity int
	rows, err := c.DB.Query("SELECT quantity FROM cart_items WHERE user_id = $1", userID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var q int
		if err := rows.Scan(&q); err != nil {
			return 0, err
		}
		totalQuantity += q
	}
	return totalQuantity, nil
}
