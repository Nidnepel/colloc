package handler

import (
	"SoftDesignColloc/internal/model"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createReview(c *gin.Context) {
	var input model.Review
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	id, err := h.services.Review.Create(ctx, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getReviewsForHotel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("hotelId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	item, err := h.services.Review.ReadReviewByHotelID(ctx, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}
