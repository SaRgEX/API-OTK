package handler

import (
	"net/http"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/gin-gonic/gin"
)

type getAllManufacturerResponse struct {
	Data []model.Manufacturer `json:"data"`
}

func (h *Handler) findAllManufacturer(c *gin.Context) {
	output, err := h.services.Manufacturer.FindAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllManufacturerResponse{
		Data: output,
	})
}

func (h *Handler) createManufacturer(c *gin.Context) {
	var input model.Manufacturer

	id, err := h.services.Manufacturer.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"name": id,
	})
}
