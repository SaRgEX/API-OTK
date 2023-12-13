package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.Account) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
	Logout(token string) error
}

type Product interface {
	Create(product model.Product) (int, error)
	FindAll() ([]model.ProductsOutput, error)
	FindById(id int) (model.Product, error)
	Update(id int, input model.UpdateProductInput) error
	Delete(id int) error
}

type Address interface {
	Create(address model.Address) (int, error)
}

type Order interface {
	Create(model.CreateInputOrder) (int, error)
	View(accountId int) ([]model.Order, error)
}

type Purchase interface {
}

type Category interface {
	FindAll() ([]model.Category, error)
	Create(model.Category) (int, error)
}

type Manufacturer interface {
	FindAll() ([]model.Manufacturer, error)
	Create(model.Manufacturer) (int, error)
}

type Service struct {
	Authorization
	Product
	Address
	Order
	Purchase
	Category
	Manufacturer
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Product:       NewProductService(repos.Product),
		Address:       NewAddressService(repos.Address),
		Order:         NewOrderService(repos.Order),
		Category:      NewCategoryService(repos.Category),
		Manufacturer:  NewManufacturerService(repos.Manufacturer),
	}
}
