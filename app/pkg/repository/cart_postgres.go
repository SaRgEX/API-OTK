package repository

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type CartPostgres struct {
	db *sqlx.DB
}

func NewCartPostgres(db *sqlx.DB) *CartPostgres {
	return &CartPostgres{db: db}
}

func (r *CartPostgres) ViewCart(userId int) (*model.CartOutput, error) {
	var output model.CartOutput
	err := r.db.Get(&output.Id, `SELECT id FROM cart WHERE account_id = $1`, userId)
	if err != nil {
		return nil, err
	}
	err = r.db.Select(&output.Products, `SELECT article, name, category, manufacturer, price, image, description, pc.amount FROM product 
	JOIN product_cart as pc ON product.article = pc.product_article
	WHERE cart_id = $1`, output.Id)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

func (r *CartPostgres) AddProduct(props model.AddToCartProps) error {
	var cart int
	err := r.db.Get(&cart, `SELECT id FROM cart WHERE account_id = $1`, props.Account)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`INSERT INTO product_cart (cart_id, product_article, amount) VALUES ($1, $2, $3)`, cart, props.Product, props.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r *CartPostgres) RemoveProduct(product int, cart int) error {
	_, err := r.db.Exec(`DELETE FROM product_cart WHERE product_article = $1 AND cart_id = $2`, product, cart)
	if err != nil {
		return err
	}

	return nil
}

func (r *CartPostgres) RemoveCart(userId int) error {
	_, err := r.db.Exec(
		`DELETE FROM product_cart
		USING cart
		WHERE cart.id = product_cart.cart_id AND cart.account_id = $1`,
		userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *CartPostgres) ScaleProduct(input model.ScaleProductInput) error {
	_, err := r.db.Exec(`UPDATE product_cart SET amount = $1 WHERE product_article = $2 AND cart_id = $3`, input.Amount, input.Product, input.Cart)
	if err != nil {
		return err
	}

	return nil
}
