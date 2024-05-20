package handler

import (
	"net/http"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/gin-gonic/gin"
)

// @Summary Create address
// @Tags address
// @Security ApiKeyAuth
// @Description create address
// @ID create-address
// @Accept  json
// @Produce  json
// @Param input body model.Address true "address info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/address [post]
func (h *Handler) createAddress(c *gin.Context) {
	var input model.ClientAddressInput

	accountId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.ClientId = accountId

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Address.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get all address
// @Tags address
// @Security ApiKeyAuth
// @Description get all address
// @ID get-all-address
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Address
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/address [get]
func (h *Handler) getAllAddress(c *gin.Context) {
	accountId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	address, err := h.services.Address.GetAll(accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, address)
}
