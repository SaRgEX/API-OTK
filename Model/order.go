package model

type Order struct {
	Id        int    `json:"-"`
	Address   int    `json:"address"`
	OrderDate string `json:"order_date"`
	AccountId int    `json:"account_id"`
	Status    string `json:"status"`
}
