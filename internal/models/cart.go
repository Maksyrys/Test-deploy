package models

import "time"

type CartItem struct {
	CartItemID int       `json:"cart_item_id"`
	UserID     int       `json:"user_id"`
	BookID     int       `json:"book_id"`
	Quantity   int       `json:"quantity"`
	AddedAt    time.Time `json:"added_at"`
}
