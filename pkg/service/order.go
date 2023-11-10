package service

import (
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(accountId, addressId, productArticle, amount int, order_date, status string) (int, error) {
	return s.repo.Create(accountId, addressId, productArticle, amount, order_date, status)
}
