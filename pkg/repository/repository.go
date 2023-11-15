package repository

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.Account) (int, error)
	GetUser(login, password string) (model.Account, error)
}

type Product interface {
	Create(product model.Product) (int, error)
	FindAll() ([]model.Product, error)
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

type Repository struct {
	Authorization
	Product
	Address
	Order
	Purchase
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Product:       NewProductPostgres(db),
		Address:       NewAddressPostgres(db),
		Order:         NewOrderPostgres(db),
	}
}