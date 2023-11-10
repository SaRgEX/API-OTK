package model

type Product struct {
	Article      int    `json:"article"`
	Name         string `json:"name" binding:"required"`
	Category     string `json:"category"`
	Manufacturer string `json:"manufacturer"`
	Price        int    `json:"price"`
	Image        string `json:"image"`
	Description  string `json:"description"`
}
