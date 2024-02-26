package repository

import "github.com/jmoiron/sqlx"

type FavoritePostgres struct {
	db *sqlx.DB
}

func NewFavoritePostgres(db *sqlx.DB) *FavoritePostgres {
	return &FavoritePostgres{db: db}
}

func (r *FavoritePostgres) AddToFavorite(userId int, productId int) error {
	_, err := r.db.Exec(`INSERT INTO favorite (account_id, product_id) VALUES ($1, $2)`, userId, productId)
	return err
}

func (r *FavoritePostgres) FindAll(userId int) ([]int, error) {
	var products []int
	err := r.db.Select(&products, `SELECT product_id FROM favorite WHERE account_id = $1`, userId)
	return products, err
}
