package handler

import (
	"SoftDesignColloc/internal/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	hotels := router.Group("/hotels")
	{
		hotels.GET("/", h.getAllHotels)
		hotels.GET("/:id", h.getHotelByID)
	}

	bookings := router.Group("/bookings")
	{
		bookings.POST("/", h.createBooking)
	}

	reviews := router.Group("/reviews")
	{
		reviews.POST("/", h.createReview)
		reviews.GET("/:hotelId", h.getReviewsForHotel)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
