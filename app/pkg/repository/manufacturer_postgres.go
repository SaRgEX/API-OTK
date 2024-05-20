package repository

import (
	"fmt"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type ManufacturerPostgres struct {
	db *sqlx.DB
}

func NewManufacturerPostgres(db *sqlx.DB) *ManufacturerPostgres {
	return &ManufacturerPostgres{db: db}
}

func (r *ManufacturerPostgres) FindAll() ([]model.Manufacturer, error) {
	var manufacturers []model.Manufacturer

	query := fmt.Sprintf(`SELECT name FROM %s`, manufacturerTable)

	err := r.db.Select(&manufacturers, query)

	return manufacturers, err
}

func (r *ManufacturerPostgres) Create(input model.Manufacturer) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO manufacturer(name) VALUES ($1)", input.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
