package handler

import (
	"SoftDesignColloc/internal/model"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary create Booking
// @Description create Booking
// @Produce json
// @Router /bookings [post]
func (h *Handler) createBooking(c *gin.Context) {
	var input model.Booking
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	id, err := h.services.Booking.Create(ctx, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
