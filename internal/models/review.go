package models

import "time"

type Review struct {
	ReviewID  int
	UserID    int
	BookID    int
	Username  string // Имя пользователя для отображения
	BookTitle string // Название книги для удобства
	Rating    int
	Comment   string
	Created   time.Time
}
