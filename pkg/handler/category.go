package handler

import (
	"net/http"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/gin-gonic/gin"
)

type getAllCategoryResponse struct {
	Data []model.Category `json:"data"`
}

func (h *Handler) findAllCategory(c *gin.Context) {
	output, err := h.services.Category.FindAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCategoryResponse{
		Data: output,
	})
}

func (h *Handler) createCategory(c *gin.Context) {
	var input model.Category

	id, err := h.services.Category.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"name": id,
	})
}
