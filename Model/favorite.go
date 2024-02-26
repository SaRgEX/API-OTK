package model

type Favorite struct {
	ProductId int `json:"product_id" db:"product_id"`
	UserId    int `json:"account_id" db:"account_id"`
}

type FavoritesOutput struct {
	Product []ProductsOutput `json:"product"`
}
