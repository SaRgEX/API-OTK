package repository

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type FavoritePostgres struct {
	db *sqlx.DB
}

func NewFavoritePostgres(db *sqlx.DB) *FavoritePostgres {
	return &FavoritePostgres{db: db}
}

func (r *FavoritePostgres) AddToFavorite(userId int, productId int) error {
	_, err := r.db.Exec(`INSERT INTO favorite (account_id, product_article) VALUES ($1, $2)`, userId, productId)
	return err
}

func (r *FavoritePostgres) FindAll(userId int) (*model.FavoritesOutput, error) {
	var products model.FavoritesOutput
	err := r.db.Select(&products.Product, `SELECT article, name, category, manufacturer, price, image, description FROM product
		JOIN favorite ON article = product_article
		JOIN account ON account_id = $1
		GROUP BY article`, userId)
	return &products, err
}

func (r *FavoritePostgres) RemoveFromFavorite(userId int, productId int) error {
	_, err := r.db.Exec(`DELETE FROM favorite WHERE product_article = $1 AND account_id = $2`, productId, userId)
	if err != nil {
		return err
	}

	return nil
}
