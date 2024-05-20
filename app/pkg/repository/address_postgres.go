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

func (r *AddressPostgres) Create(address model.ClientAddressInput) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("INSERT INTO %s (street, house, apartment, postal_code, city) VALUES ($1, $2, $3, $4, $5) RETURNING id", addressTable)

	row := tx.QueryRow(query, address.Address.Street, address.Address.House, address.Address.Apartment, address.Address.PostalCode, address.Address.City)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	query = fmt.Sprintf("INSERT INTO %s (account_id, address_id) VALUES ($1, $2)", clientAddressTable)
	_, err = tx.Exec(query, address.ClientId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *AddressPostgres) GetAll(accountId int) (*model.ClientAddressOutput, error) {
	var addresses model.ClientAddressOutput
	query := fmt.Sprintf(
		`SELECT a.id, street, house, apartment, postal_code, city FROM %s AS a
		JOIN %s AS ca ON ca.address_id = a.id
		WHERE ca.account_id = $1
		GROUP BY id`,
		addressTable, clientAddressTable)
	err := r.db.Select(&addresses.Address, query, accountId)
	return &addresses, err
}
