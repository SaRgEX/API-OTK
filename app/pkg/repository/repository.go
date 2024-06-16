package repository

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	CreateUserWithRole(user model.UserWithRole) (int, error)
	GetUser(login, password string) (model.UserWithRole, error)
	Logout(token string) error
}

type User interface {
	Find(id int) (model.UserOutput, error)
	UpdateUser(id int, input model.UpdateUser) error
}

type Product interface {
	Create(product model.Product) (int, error)
	FindAll() ([]model.ProductsOutput, error)
	FindById(id int) (model.ProductsOutput, error)
	Update(id int, input model.UpdateProductInput) (int, error)
	Delete(id int) error
}

type Address interface {
	Create(address model.ClientAddressInput) (int, error)
	GetAll(accountId int) (*model.ClientAddressOutput, error)
}

type Order interface {
	Create(model.CreateInputOrder) (int, error)
	View(accountId int) ([]model.OrderOutput, error)
	ViewOne(id, account_id int) (model.OrderOutputProps, error)
	AdminOrder() ([]model.AdminOrderOutput, error)
	UpdateOrderStatus(id int, status model.UpdateOrderStatus) (int, error)
	OrderStatus() ([]model.OrderStatusOutputs, error)
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
	RemoveCart(cart int) error
}

type Favorite interface {
	AddToFavorite(userId int, productId int) error
	FindAll(userId int) (*model.FavoritesOutput, error)
	RemoveFromFavorite(userId int, productId int) error
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
