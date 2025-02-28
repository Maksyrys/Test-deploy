package models

import "time"

type Book struct {
	ID                  int       `json:"id"`
	Title               string    `json:"title"`
	Author              string    `json:"author"`
	Category            string    `json:"category"`
	Price               float64   `json:"price"`
	Description         string    `json:"description"`          // краткое описание
	DetailedDescription string    `json:"detailed_description"` // новое поле
	PublishDate         time.Time `json:"publish_date"`
	ImageURL            string    `json:"image_url"`
}
