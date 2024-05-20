package model

type Manufacturer struct {
	Name string `json:"manufacturer" binding:"required" db:"name"`
}
