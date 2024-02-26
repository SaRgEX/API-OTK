package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}
func (s *CartService) ViewCart(userId int) (*model.CartOutput, error) {
	return s.repo.ViewCart(userId)
}

func (s *CartService) AddProduct(props model.AddToCartProps) error {
	return s.repo.AddProduct(props)
}

func (s *CartService) RemoveProduct(product int, cart int) error {
	return s.repo.RemoveProduct(product, cart)
}

func (s *CartService) ScaleProduct(input model.ScaleProductInput) error {
	return s.repo.ScaleProduct(input)
}
