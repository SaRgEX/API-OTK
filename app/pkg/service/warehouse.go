package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type WarehouseService struct {
	repo repository.Warehouse
}

func NewWarehouseService(repo repository.Warehouse) *WarehouseService {
	return &WarehouseService{repo: repo}
}

func (s *WarehouseService) FindAll() ([]model.Warehouse, error) {
	return s.repo.FindAll()
}
