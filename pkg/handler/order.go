package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ProductArticle int    `json:"product_article" binding:"required"`
	Amount         int    `json:"amount" binding:"required"`
	Address        int    `json:"address" binding:"required"`
	OrderDate      string `json:"order_date" binding:"required"`
	Status         string `json:"status" binding:"required"`
}

func (h *Handler) createOrder(c *gin.Context) {
	var input Order

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accountId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Printf("account_id type: %v, address_id type: %v, product_article type: %v, amount type: %v", accountId, input.Address, input.ProductArticle, input.Amount)
	id, err := h.services.Order.Create(accountId, input.Address, input.ProductArticle, input.Amount, input.OrderDate, input.Status)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
