package handler

import (
	"net/http"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createOrder(c *gin.Context) {
	var input model.CreateInputOrder

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
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

func (h *Handler) viewOrders(c *gin.Context) {
	var accountId int

	accountId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input, err := h.services.Order.View(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	for _, order := range input {
		c.JSON(http.StatusOK, map[string]interface{}{
			"address":    order.Address,
			"order_date": order.OrderDate,
			"account_id": order.AccountId,
			"status":     order.Status,
		})
	}
}
