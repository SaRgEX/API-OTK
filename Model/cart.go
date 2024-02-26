package model

type AddToCartInput struct {
	Product int `json:"product_article" db:"product_article"`
	Amount  int `json:"amount" db:"amount"`
}

type AddToCartProps struct {
	AddToCartInput
	Account int `json:"account_id" db:"account_id"`
}

type Cart struct {
	Id       int       `json:"id" db:"id"`
	Products []Product `json:"product_article" db:"product_article"`
	UserId   int       `json:"account_id" db:"account_id"`
}

type CartOutput struct {
	Id       int              `json:"id"`
	Products []ProductsOutput `json:"products"`
}

type RemoveProductInput struct {
	Product int `json:"product_article" db:"product_article"`
	Cart    int `json:"cart_id" db:"cart_id"`
}

type ScaleProductInput struct {
	Product int `json:"product_article" db:"product_article" binding:"required"`
	Amount  int `json:"amount" db:"amount" binding:"required"`
	Cart    int `json:"cart_id" db:"cart_id" binding:"required"`
}
