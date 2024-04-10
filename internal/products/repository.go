package products

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var products []Product
var masterId uint = 0

type Repository interface {
	GetAll() ([]Product, error)
	Create(p Product) (Product, error)
	Delete(id uint) error
	Update(id uint, p Product) (Product, error)
	UpdateFull(p Product) (Product, error)
	LastId() (uint, error)
}

type repository struct{}

func ReadFile() error {
	jsonFile, err := os.Open(`products.json`)
	if err != nil {
		return err
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &products)
	for _, p := range products {
		if masterId < p.Id {
			masterId = p.Id
		}
	}
	defer jsonFile.Close()
	return nil
}

func UpdateFile() error {
	//abrir o arquivo products.json
	jsonFile, err := os.OpenFile("products.json", os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	//transformar o products em json
	productsJson, err := json.MarshalIndent(products, "", "\t")
	if err != nil {
		return err
	}
	//limpar o arquivo products.json
	err = os.Truncate("products.json", 0)
	if err != nil {
		return err
	}
	//escrever no arquivo
	_, err = jsonFile.Write(productsJson)
	if err != nil {
		return fmt.Errorf("erro no metodo write: %w", err)
	}
	return nil
}

func (r *repository) GetAll() ([]Product, error) {
	if len(products) == 0 {
		err := ReadFile()
		if err != nil {
			return []Product{}, err
		}
	}
	return products, nil
}

func (r *repository) Create(p Product) (Product, error) {
	if len(products) == 0 {
		err := ReadFile()
		if err != nil {
			return Product{}, err
		}
	}
	masterId++
	p.Id = masterId
	products = append(products, p)
	err := UpdateFile()
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Delete(id uint) error {
	if len(products) == 0 {
		err := ReadFile()
		if err != nil {
			return err
		}
	}
	for i, prod := range products {
		if prod.Id == id {
			products = append(products[:i], products[i+1:]...)
		}
	}
	err := UpdateFile()
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateFull(p Product) (Product, error) {
	updated := false
	if len(products) == 0 {
		err := ReadFile()
		if err != nil {
			return Product{}, err
		}
	}
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
	err := UpdateFile()
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Update(id uint, p Product) (Product, error) {
	if len(products) == 0 {
		err := ReadFile()
		if err != nil {
			return Product{}, err
		}
	}
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
	err := UpdateFile()
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) LastId() (uint, error) {
	return 0, nil
}

func NewRepository() Repository {
	return &repository{}
}
