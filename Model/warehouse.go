package model

type Warehouse struct {
	Id      int `json:"id" db:"id"`
	Address int `json:"address" db:"address"`
}
