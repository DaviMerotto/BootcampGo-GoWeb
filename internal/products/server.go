package products

type Service interface {
	GetAll() ([]Product, error)
	Create(p Product) (Product, error)
	Delete(id uint) error
	UpdateFull(p Product) (Product, error)
	Update(id uint, p Product) (Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	prods, err := s.repository.GetAll()
	if err != nil {
		return []Product{}, err
	}
	return prods, nil
}

func (s *service) Create(p Product) (Product, error) {
	product, err := s.repository.Create(p)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Delete(id uint) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateFull(p Product) (Product, error) {
	product, err := s.repository.UpdateFull(p)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Update(id uint, p Product) (Product, error) {
	product, err := s.repository.Update(id, p)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
