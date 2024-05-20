package handler

import (
	"net/http"
	"strconv"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/gin-gonic/gin"
)

// @Summary Create order
// @Tags order
// @Security ApiKeyAuth
// @Description create order
// @ID create-order
// @Accept  json
// @Produce  json
// @Param input body model.CreateInputOrder true "order info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/order [post]
func (h *Handler) createOrder(c *gin.Context) {
	var input model.CreateInputOrder

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(input.Purchase) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "empty order")
		return
	}

	accountId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.AccountId = accountId

	id, err := h.services.Order.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type OrdeData struct {
	Data []model.Order `json:"data"`
}

// @Summary View orders
// @Tags order
// @Security ApiKeyAuth
// @Description view orders
// @ID view-orders
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/order [get]
func (h *Handler) viewOrders(c *gin.Context) {
	var accountId int

	accountId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	output, err := h.services.Order.View(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, &output)

}

func (h *Handler) viewOrder(c *gin.Context) {
	var accountId int
	accountId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	output, err := h.services.Order.ViewOne(id, accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, &output)
}
