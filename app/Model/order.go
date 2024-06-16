package model

type Order struct {
	Id        int    `json:"-"`
	Address   int    `json:"address" db:"address"`
	OrderDate string `json:"order_date" db:"order_date"`
	AccountId int    `json:"account_id" db:"account_id"`
	Status    string `json:"status" db:"status"`
}

type Purchase struct {
	ProductArticle int `json:"product_article" binding:"required"`
	Amount         int `json:"amount" binding:"required"`
}
type CreateInputOrder struct {
	AccountId int        `json:"-"`
	Purchase  []Purchase `json:"purchase"`
	Address   int        `json:"address" binding:"required"`
	OrderDate string     `json:"-"`
	Status    string     `json:"-"`
}

type OrderOutputProps struct {
	OrderOutput OrderOutput      `json:"order"`
	Purchase    []PurchaseOutput `json:"purchase"`
}

type OrderOutput struct {
	Id        int    `json:"id" db:"id"`
	OrderDate string `json:"order_date" db:"order_date"`
	Status    string `json:"status" db:"status"`

	Street    string `json:"street" db:"street"`
	House     string `json:"house" db:"house"`
	Apartment string `json:"apartment" db:"apartment"`
	City      string `json:"city" db:"city"`
}

type AdminOrderOutput struct {
	OrderOutput
	AccountId int `json:"account_id" db:"account_id"`
	UserOutputAdmin
}

type OrderStatusOutputs struct {
	Status string `json:"status" db:"name"`
}

type PurchaseOutput struct {
	Product
	Amount int `json:"amount" db:"amount"`
}

type UpdateOrderStatus struct {
	Status string `json:"order_status" binding:"required"`
}
