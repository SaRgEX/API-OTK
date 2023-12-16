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
	_, err = tx.Exec(createPurchaseQuery, model.ProductArticle, id, model.Amount)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *OrderPostgres) View(accountId int) ([]model.OrderOutput, error) {
	var orderModel []model.OrderOutput
	getOrderQuery := fmt.Sprintf(
		`SELECT order_date, status, a.street, a.house, a.apartment, a.city FROM %s
		JOIN address a ON address = a.id
		WHERE account_id = $1`,
		orderTable)

	err := r.db.Select(&orderModel, getOrderQuery, accountId)
	return orderModel, err
}
