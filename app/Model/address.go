package model

type Address struct {
	Id         int    `json:"id" db:"id"`
	Street     string `json:"street" db:"street" binding:"required"`
	House      string `json:"house" db:"house" binding:"required"`
	Apartment  string `json:"apartment" db:"apartment" binding:"required"`
	PostalCode int    `json:"postal_code" db:"postal_code" binding:"required"`
	City       string `json:"city" db:"city" binding:"required"`
}

type ClientAddressOutput struct {
	ClientId int       `json:"-"`
	Address  []Address `json:"address"`
}

type ClientAddressInput struct {
	ClientId int `json:"-"`
	Address
}
