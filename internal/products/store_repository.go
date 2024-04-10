package products

import (
	"fmt"

	"github.com/davimerotto/web-server/pkg/store"
)

type StoreFileRepository struct {
	db store.Store
}

func NewStoreRepository(db store.Store) Repository {
	return &StoreFileRepository{
		db: db,
	}
}

func (r *StoreFileRepository) GetAll() ([]Product, error) {
	var products []Product
	r.db.Read(&products)
	return products, nil
}

func (r *StoreFileRepository) Create(p Product) (Product, error) {
	var products []Product

	r.db.Read(&products)

	lastIdInserted, err := r.LastId()
	if err != nil {
		return Product{}, err
	}
	lastIdInserted++
	p.Id = uint(lastIdInserted)

	products = append(products, p)

	err = r.db.Write(products)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *StoreFileRepository) Delete(id uint) error {
	var products []Product
	r.db.Read(&products)

	for i, prod := range products {
		if prod.Id == id {
			products = append(products[:i], products[i+1:]...)
		}
	}
	err := r.db.Write(products)
	if err != nil {
		return err
	}
	return nil
}

func (r *StoreFileRepository) UpdateFull(p Product) (Product, error) {
	updated := false
	var products []Product
	r.db.Read(&products)

	for i, prod := range products {
		if prod.Id == p.Id {
			products[i] = p
			p = products[i]
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Product not found")
	}
	err := r.db.Write(products)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *StoreFileRepository) Update(id uint, p Product) (Product, error) {
	var products []Product
	r.db.Read(&products)
	updated := false

	for i, prod := range products {
		if prod.Id == id {
			if p.Name != "" {
				products[i].Name = p.Name
			}
			if p.Color != "" {
				products[i].Color = p.Color
			}
			if p.Price != 0 {
				products[i].Price = p.Price
			}
			if p.Stock != 0 {
				products[i].Stock = p.Stock
			}
			if p.Code != "" {
				products[i].Code = p.Code
			}
			if p.Published {
				products[i].Published = p.Published
			}
			if p.Creation_date != "" {
				products[i].Creation_date = p.Creation_date
			}
			updated = true
			p = products[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Product not found")
	}
	err := r.db.Write(products)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *StoreFileRepository) LastId() (uint, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}
