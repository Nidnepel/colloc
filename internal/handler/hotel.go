package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getAllHotels(c *gin.Context) {
	items, err := h.services.Hotel.ReadAll(context.Background())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getHotelByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	item, err := h.services.Hotel.Read(ctx, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}
