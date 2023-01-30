package book

import "time"

type Book struct {
	ID          int32
	Title       string
	Description string
	Price       int32
	Rating      int32
	Discount 	int32
	CreatedAt   time.Time
	UpdatedAt 	time.Time
}