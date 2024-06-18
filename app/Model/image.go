package model

type ImageInput struct {
	Image     []byte `json:"image" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Extension string `json:"extension" binding:"required"`
}
