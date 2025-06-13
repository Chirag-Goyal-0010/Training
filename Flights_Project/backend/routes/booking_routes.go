package routes

import (
	"flights-project/controllers"
	"flights-project/middleware"

	"github.com/gin-gonic/gin"
)

func SetupBookingRoutes(router *gin.Engine, bookingController *controllers.BookingController) {
	bookingRoutes := router.Group("/api/bookings")
	bookingRoutes.Use(middleware.AuthMiddleware())

	bookingRoutes.POST("/", bookingController.CreateBooking)
	bookingRoutes.GET("/", bookingController.GetUserBookings)
	bookingRoutes.POST("/:id/cancel", bookingController.CancelBooking)
}
