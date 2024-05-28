package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(model model.CreateInputOrder) (int, error) {
	return s.repo.Create(model)
}

func (s *OrderService) View(accountId int) ([]model.OrderOutput, error) {
	return s.repo.View(accountId)
}

func (s *OrderService) ViewOne(id, accountId int) (model.OrderOutputProps, error) {
	return s.repo.ViewOne(id, accountId)
}
