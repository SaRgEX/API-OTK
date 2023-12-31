package model

type Order struct {
	Id        int    `json:"-"`
	Address   int    `json:"address" db:"address"`
	OrderDate string `json:"order_date" db:"order_date"`
	AccountId int    `json:"account_id" db:"account_id"`
	Status    string `json:"status" db:"status"`
}

type CreateInputOrder struct {
	AccountId      int    `json:"-"`
	ProductArticle int    `json:"product_article" binding:"required"`
	Amount         int    `json:"amount" binding:"required"`
	Address        int    `json:"address" binding:"required"`
	OrderDate      string `json:"-"`
	Status         string `json:"-"`
}

type OrderOutput struct {
	OrderDate string `json:"order_date" db:"order_date"`
	Status    string `json:"status" db:"status"`

	Street    string `json:"street" db:"street"`
	House     string `json:"house" db:"house"`
	Apartment string `json:"apartment" db:"apartment"`
	City      string `json:"city" db:"city"`
}
