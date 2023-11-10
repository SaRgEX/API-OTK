package repository

import (
	"fmt"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type AddressPostgres struct {
	db *sqlx.DB
}

func NewAddressPostgres(db *sqlx.DB) *AddressPostgres {
	return &AddressPostgres{db: db}
}

func (r *AddressPostgres) Create(address model.Address) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (street, house, apartment, postal_code, city) VALUES ($1, $2, $3, $4, $5) returning id",
		addressTable)
	row := r.db.QueryRow(
		query,
		address.Street,
		address.House,
		address.Apartment,
		address.PostalCode,
		address.City)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
