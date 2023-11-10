package repository

import (
	"fmt"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product model.Product) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (article, name, category, manufacturer, price, image, description) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING article", productTable)
	row := r.db.QueryRow(
		query,
		product.Article,
		product.Name,
		product.Category,
		product.Manufacturer,
		product.Price,
		product.Image,
		product.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
