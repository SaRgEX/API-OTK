package repository

import (
	"errors"
	"fmt"
	"strings"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r *ProductPostgres) FindAll() ([]model.ProductsOutput, error) {
	var products []model.ProductsOutput
	query := fmt.Sprintf(`SELECT article, name, category, manufacturer, price, image, description, SUM(ps.amount) as amount FROM %s
	INNER JOIN %s AS ps ON article = ps.product_article
	GROUP BY article`, productTable, productStackTable)
	err := r.db.Select(&products, query)
	return products, err
}

func (r *ProductPostgres) FindById(article int) (model.ProductsOutput, error) {
	var product model.ProductsOutput
	query := fmt.Sprintf(`SELECT article, name, category, manufacturer, price, image, description, SUM(ps.amount) as amount FROM %s 
	INNER JOIN %s AS ps ON article = ps.product_article
	WHERE article = $1
	GROUP BY article`, productTable, productStackTable)
	err := r.db.Get(&product, query, article)
	return product, err
}

func (r *ProductPostgres) Delete(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE product_article = $1", productStackTable)
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE article = $1 RETURNING name, category, manufacturer, price, image, description", productTable)
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *ProductPostgres) Update(id int, input model.UpdateProductInput) (int, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	ardId := 1

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", ardId))
		args = append(args, *input.Name)
		ardId++
	}

	if input.Category != nil {
		setValues = append(setValues, fmt.Sprintf("category=$%d", ardId))
		args = append(args, *input.Category)
		ardId++
	}

	if input.Manufacturer != nil {
		setValues = append(setValues, fmt.Sprintf("manufacturer=$%d", ardId))
		args = append(args, *input.Manufacturer)
		ardId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", ardId))
		args = append(args, *input.Price)
		ardId++
	}

	if input.Image != nil {
		setValues = append(setValues, fmt.Sprintf("image=$%d", ardId))
		args = append(args, *input.Image)
		ardId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", ardId))
		args = append(args, *input.Description)
		ardId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE article=$%d", productTable, setQuery, ardId)

	args = append(args, id)

	logrus.Debugf("update query: %s", query)
	logrus.Debugf("update args: %s", args)

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if (input.Amount == nil) && (input.Warehouse == nil) {
		return id, tx.Commit()
	} else if (input.Amount == nil) || (input.Warehouse == nil) {
		tx.Rollback()
		return 0, errors.New("amount or warehouse is empty")
	}
	query = fmt.Sprintf("UPDATE %s SET amount = $1 WHERE product_article = $2 AND warehouse_id = $3", productStackTable)
	_, err = tx.Exec(query, *input.Amount, id, *input.Warehouse)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
