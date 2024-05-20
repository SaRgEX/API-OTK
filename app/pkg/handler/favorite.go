package handler

import (
	"net/http"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/gin-gonic/gin"
)

// @Summary Add to favorite
// @Tags favorite
// @Description add product to favorite
// @ID add-to-favorite
// @Accept  json
// @Produce  json
// @Param input body model.Favorite true "favorite info"
// @Success 200 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/favorite [post]
func (h *Handler) addToFavorite(c *gin.Context) {
	var input model.Favorite
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Favorite.AddToFavorite(userId, input.ProductId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// @Summary View favorites
// @Tags favorite
// @Security ApiKeyAuth
// @Description view user's products in favorite list
// @ID view-favorite
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/favorite [get]
func (h *Handler) findAllFavorite(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	output, err := h.services.Favorite.FindAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, output)
}

func (h *Handler) removeFavorite(c *gin.Context) {
	var input model.Favorite
	err := c.ShouldBind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.Favorite.RemoveFromFavorite(userId, input.ProductId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"removed"})
}
