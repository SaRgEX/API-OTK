package handler

import (
	"net/http"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/gin-gonic/gin"
)

type getAllWarehouseResponse struct {
	Data []model.Warehouse `json:"data"`
}

func (h *Handler) findAllWarehouse(c *gin.Context) {
	output, err := h.services.Warehouse.FindAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWarehouseResponse{
		Data: output,
	})
}
