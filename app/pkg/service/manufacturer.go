package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type ManufacturerService struct {
	repo repository.Manufacturer
}

func NewManufacturerService(repo repository.Manufacturer) *ManufacturerService {
	return &ManufacturerService{repo: repo}
}

func (s *ManufacturerService) FindAll() ([]model.Manufacturer, error) {
	return s.repo.FindAll()
}

func (s *ManufacturerService) Create(input model.Manufacturer) (int, error) {
	return s.repo.Create(input)
}
