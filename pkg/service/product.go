package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(product model.Product) (int, error) {
	return s.repo.Create(product)
}

func (s *ProductService) FindAll() ([]model.ProductsOutput, error) {
	return s.repo.FindAll()
}

func (s *ProductService) FindById(id int) (model.Product, error) {
	return s.repo.FindById(id)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *ProductService) Update(id int, input model.UpdateProductInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
