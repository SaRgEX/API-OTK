package handler

import (
	"net/http"
	"strconv"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/gin-gonic/gin"
)

// @Summary Create product
// @Tags product
// @Description create product
// @ID create-product
// @Accept  json
// @Produce  json
// @Param input body model.Product true "product info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /products [post]
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
	Data []model.ProductsOutput `json:"data"`
}

// @Summary Get all products
// @Tags product
// @Security ApiKeyAuth
// @Description get all products
// @ID get-all-products
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /products [get]
func (h *Handler) findAll(c *gin.Context) {
	output, err := h.services.Product.FindAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllProductsResponse{
		Data: output,
	})
}

// @Summary Get product
// @Tags product
// @Security ApiKeyAuth
// @Description get product
// @ID get-product
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /products/{id} [get]
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

// @Summary Update product
// @Tags product
// @Security ApiKeyAuth
// @Description update product
// @ID update-product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param input body model.UpdateProductInput true "Product data"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /products/{id} [put]
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

	_, err = h.services.Product.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete product
// @Tags product
// @Security ApiKeyAuth
// @Description delete product
// @ID delete-product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /products/{id} [delete]
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
