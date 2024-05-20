package repository

import (
	"fmt"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) Create(model model.CreateInputOrder) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createOrderQuery := fmt.Sprintf(
		"INSERT INTO %s (address, account_id, order_date, status) VALUES ($1, $2, DEFAULT, DEFAULT) RETURNING id",
		orderTable,
	)
	row := tx.QueryRow(createOrderQuery, model.Address, model.AccountId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createPurchaseQuery := fmt.Sprintf(
		"INSERT INTO %s (product_article, order_id, amount) VALUES ($1, $2, $3)",
		purchaseTable,
	)

	for _, product := range model.Purchase {
		_, err = tx.Exec(createPurchaseQuery, product.ProductArticle, id, product.Amount)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *OrderPostgres) Delete(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE order_id = $1", purchaseTable)
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE id = $1", orderTable)
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *OrderPostgres) View(accountId int) ([]model.OrderOutput, error) {
	var orderModel []model.OrderOutput
	getOrderQuery := fmt.Sprintf(
		`SELECT o.id, order_date, status, a.street, a.house, a.apartment, a.city FROM %s AS o
		JOIN address a ON address = a.id
		WHERE account_id = $1`,
		orderTable)

	err := r.db.Select(&orderModel, getOrderQuery, accountId)
	return orderModel, err
}

func (r *OrderPostgres) ViewOne(id, accountId int) (model.OrderOutput, error) {
	var orderModel model.OrderOutput

	getOrderQuery := fmt.Sprintf(
		`SELECT o.id, order_date, status, a.street, a.house, a.apartment, a.city FROM %s AS o
		JOIN address a ON address = a.id
		WHERE o.id = $1`,
		orderTable)

	err := r.db.Get(&orderModel, getOrderQuery, id)
	return orderModel, err
}
