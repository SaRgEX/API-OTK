package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.Account) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Product interface {
	Create(product model.Product) (int, error)
}

type Address interface {
	Create(address model.Address) (int, error)
}

type Order interface {
	Create(accountId, addressId, productArticle, amount int, order_date, status string) (int, error)
}

type Purchase interface {
}

type Service struct {
	Authorization
	Product
	Address
	Order
	Purchase
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Product:       NewProductService(repos.Product),
		Address:       NewAddressService(repos.Address),
		Order:         NewOrderService(repos.Order),
	}
}
