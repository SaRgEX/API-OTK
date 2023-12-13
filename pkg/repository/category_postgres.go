package repository

import (
	"fmt"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) FindAll() ([]model.Category, error) {
	var category []model.Category

	query := fmt.Sprintf(`SELECT name FROM %s`, categoryTable)

	err := r.db.Select(&category, query)
	return category, err
}

func (r *CategoryPostgres) Create(input model.Category) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO category VALUES ($2)", input.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
