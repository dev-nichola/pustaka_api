package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(id int) (Book, error)
	Create(bookrequest BookRequest) (Book, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	book, err := s.repository.FindAll()
	return book, err
}

func (s *service) FindById(id int) (Book, error) {
	book, err := s.repository.FindById(id)
	return book, err
}

func (s *service) Create(bookrequest BookRequest) (Book, error) {
	price, _ := bookrequest.Price.Int64()
	rating, _ := bookrequest.Rating.Int64()
	discount, _ := bookrequest.Discount.Int64()

	book := Book{
		Title:       bookrequest.Title,
		Price:       int(price),
		Description: bookrequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}
	book, err := s.repository.Create(book)
	return book, err
}
