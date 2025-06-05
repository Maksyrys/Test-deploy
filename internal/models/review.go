package models

import "time"

type Review struct {
	ReviewID  int
	UserID    int
	BookID    int
	Username  string
	BookTitle string
	Rating    int
	Comment   string
	Created   time.Time
}
