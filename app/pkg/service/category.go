package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) FindAll() ([]model.Category, error) {
	return s.repo.FindAll()
}

func (s *CategoryService) Create(input model.Category) (int, error) {
	return s.repo.Create(input)
}
