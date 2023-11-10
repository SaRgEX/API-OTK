package model

type Address struct {
	Id         int    `json:"-"`
	Street     string `json:"street" binding:"required"`
	House      string `json:"house" binding:"required"`
	Apartment  string `json:"apartment" binding:"required"`
	PostalCode int    `json:"postal_code" binding:"required"`
	City       string `json:"city" binding:"required"`
}
