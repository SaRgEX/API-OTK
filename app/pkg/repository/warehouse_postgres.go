package repository

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type WarehousePostgres struct {
	db *sqlx.DB
}

func NewWarehousePostgres(db *sqlx.DB) *WarehousePostgres {
	return &WarehousePostgres{db: db}
}

func (r *WarehousePostgres) FindAll() ([]model.Warehouse, error) {
	var output []model.Warehouse
	err := r.db.Select(&output, "SELECT * FROM warehouse")
	return output, err
}
