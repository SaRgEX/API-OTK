package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
	Logout(token string) error
}

type User interface {
	Find(id int) (model.UserOutput, error)
}

type Product interface {
	Create(product model.Product) (int, error)
	FindAll() ([]model.ProductsOutput, error)
	FindById(id int) (model.ProductsOutput, error)
	Update(id int, input model.UpdateProductInput) (int, error)
	Delete(id int) error
}

type Address interface {
	Create(address model.Address) (int, error)
}

type Order interface {
	Create(model.CreateInputOrder) (int, error)
	View(accountId int) ([]model.OrderOutput, error)
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

type Warehouse interface {
	FindAll() ([]model.Warehouse, error)
}

type Cart interface {
	ViewCart(userId int) (*model.CartOutput, error)
	AddProduct(props model.AddToCartProps) error
	RemoveProduct(product int, cart int) error
	ScaleProduct(props model.ScaleProductInput) error
}

type Favorite interface {
	AddToFavorite(userId int, productId int) error
	FindAll(userId int) ([]int, error)
}

type Service struct {
	Authorization
	User
	Product
	Address
	Order
	Purchase
	Category
	Manufacturer
	Warehouse
	Cart
	Favorite
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Product:       NewProductService(repos.Product),
		Address:       NewAddressService(repos.Address),
		Order:         NewOrderService(repos.Order),
		Category:      NewCategoryService(repos.Category),
		Manufacturer:  NewManufacturerService(repos.Manufacturer),
		Warehouse:     NewWarehouseService(repos.Warehouse),
		Cart:          NewCartService(repos.Cart),
		Favorite:      NewFavoriteService(repos.Favorite),
	}
}
