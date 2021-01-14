package entities

import "time"

type Article struct {
	ID       int
	UserID   int
	Title    string
	Content  string
	CreateAt time.Time
}
