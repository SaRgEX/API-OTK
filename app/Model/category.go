package model

type Category struct {
	Name string `json:"category" binding:"required" db:"name"`
}
