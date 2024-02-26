package repository

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(login, password string) (model.User, error)
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
	Create(input model.Category) (int, error)
}

type Manufacturer interface {
	FindAll() ([]model.Manufacturer, error)
	Create(input model.Manufacturer) (int, error)
}

type Warehouse interface {
	FindAll() ([]model.Warehouse, error)
}

type Cart interface {
	ViewCart(userId int) (*model.CartOutput, error)
	AddProduct(props model.AddToCartProps) error
	RemoveProduct(product int, cart int) error
	ScaleProduct(input model.ScaleProductInput) error
}

type Favorite interface {
	AddToFavorite(userId int, productId int) error
	FindAll(userId int) ([]int, error)
}

type Repository struct {
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Product:       NewProductPostgres(db),
		Address:       NewAddressPostgres(db),
		Order:         NewOrderPostgres(db),
		Category:      NewCategoryPostgres(db),
		Manufacturer:  NewManufacturerPostgres(db),
		Warehouse:     NewWarehousePostgres(db),
		Cart:          NewCartPostgres(db),
		Favorite:      NewFavoritePostgres(db),
	}
}
