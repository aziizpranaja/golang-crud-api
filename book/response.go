package book

type BookResponse struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int32  `json:"price"`
	Rating      int32  `json:"rating"`
	Discount    int32  `json:"discount"`
}