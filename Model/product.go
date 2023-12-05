package model

import "errors"

type Product struct {
	Article      int    `json:"article" binding:"required" db:"article"`
	Name         string `json:"name" binding:"required" db:"name"`
	Category     string `json:"category" db:"category"`
	Manufacturer string `json:"manufacturer" db:"manufacturer"`
	Price        int    `json:"price" db:"price"`
	Image        string `json:"image" db:"image"`
	Description  string `json:"description" db:"description"`
}

type ProductsOutput struct {
	Article      int    `json:"article" binding:"required" db:"article"`
	Name         string `json:"name" binding:"required" db:"name"`
	Category     string `json:"category" db:"category"`
	Manufacturer string `json:"manufacturer" db:"manufacturer"`
	Price        int    `json:"price" db:"price"`
	Image        string `json:"image" db:"image"`
	Description  string `json:"description" db:"description"`
	Amount       int    `json:"amount" db:"amount"`
}

type UpdateProductInput struct {
	Name         *string `json:"name"`
	Category     *string `json:"category"`
	Manufacturer *string `json:"manufacturer"`
	Price        *int    `json:"price"`
	Image        *string `json:"image"`
	Description  *string `json:"description"`
}

func (p *UpdateProductInput) Validate() error {
	if p.Name == nil && p.Category == nil && p.Manufacturer == nil && p.Price == nil && p.Image == nil && p.Description == nil {
		return errors.New("nothing to update")
	}
	return nil
}
