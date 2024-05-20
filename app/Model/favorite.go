package model

type Favorite struct {
	UserId    int `json:"-"`
	ProductId int `json:"product_article" db:"product_article" binding:"required"`
}

type FavoritesOutput struct {
	Product []ProductsOutput `json:"product"`
}
