package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) Create(accountId, addressId, productArticle, amount int, order_date, status string) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	createOrderQuery := fmt.Sprintf(
		"INSERT INTO %s (address, order_date, account_id, status) VALUES ($1, $2, $3, $4) RETURNING id",
		orderTable,
	)
	row := tx.QueryRow(createOrderQuery, addressId, order_date, accountId, status)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createPurchaseQuery := fmt.Sprintf(
		"INSERT INTO %s (product_article, order_id, amount) VALUES ($1, $2, $3)",
		purchaseTable,
	)
	_, err = tx.Exec(createPurchaseQuery, productArticle, id, amount)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
