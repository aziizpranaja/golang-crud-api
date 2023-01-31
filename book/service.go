package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	// Cara return pertama
	books, err := s.repository.FindAll()
	return books, err

	// Cara return ke dua
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Price:       int32(price),
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
		Discount:    bookRequest.Discount,
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}