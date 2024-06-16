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

func (r *OrderPostgres) ViewOne(id, accountId int) (model.OrderOutputProps, error) {
	var orderModel model.OrderOutputProps

	getOrderQuery := fmt.Sprintf(
		`SELECT o.id, order_date, status, a.street, a.house, a.apartment, a.city FROM %s AS o
		JOIN address a ON address = a.id
		WHERE o.id = $1`,
		orderTable)

	err := r.db.Get(&orderModel.OrderOutput, getOrderQuery, id)

	if err != nil {
		return orderModel, err
	}

	getPurchaseQuery := fmt.Sprintf(
		`SELECT p.article, p.name, p.manufacturer, p.category, p.price, p.image, p.description, purchase.amount, p.price FROM %s
		JOIN %s AS p ON product_article = p.article
		WHERE order_id = $1`,
		purchaseTable,
		productTable,
	)

	err = r.db.Select(&orderModel.Purchase, getPurchaseQuery, id)

	return orderModel, err
}

func (r *OrderPostgres) AdminOrder() ([]model.AdminOrderOutput, error) {
	var orderModel []model.AdminOrderOutput
	query := fmt.Sprintf(
		`SELECT o.id, order_date, o.status, a.city, a.street, a.house, a.apartment, account_id, ac.first_name, ac.last_name, ac.patronumic, ac.phone, ac.email FROM %s AS o
		JOIN %s a ON address = a.id
		JOIN %s AS ac ON account_id = ac.id`,
		orderTable, addressTable, userTable,
	)

	err := r.db.Select(&orderModel, query)
	return orderModel, err
}

func (r *OrderPostgres) UpdateOrderStatus(id int, status model.UpdateOrderStatus) (int, error) {
	query := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", orderTable)
	_, err := r.db.Exec(query, status.Status, id)
	return id, err
}

func (r *OrderPostgres) OrderStatus() ([]model.OrderStatusOutputs, error) {
	var orderStatus []model.OrderStatusOutputs

	query := fmt.Sprintf(
		`SELECT name FROM %s`,
		orderStatusTable,
	)
	err := r.db.Select(&orderStatus, query)

	return orderStatus, err
}
