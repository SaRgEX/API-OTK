package handler

import (
	"net/http"
	"strconv"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct(c *gin.Context) {
	var input model.Product

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Product.Create(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"article": id,
	})

}

type getAllProductsResponse struct {
	Data []model.Product `json:"data"`
}

func (h *Handler) findAll(c *gin.Context) {
	input, err := h.services.Product.FindAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllProductsResponse{
		Data: input,
	})
}

func (h *Handler) findById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	product, err := h.services.Product.FindById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.UpdateProductInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Product.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Product.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusNoContent, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
