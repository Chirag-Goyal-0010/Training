package routes

import (
	"flights-project/controllers"
	"flights-project/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize controllers
	authController := controllers.NewAuthController(db)
	flightController := controllers.NewFlightController(db)
	bookingController := controllers.NewBookingController(db)

	// Public routes
	public := router.Group("/api")
	{
		// Auth routes
		public.POST("/auth/register",
			middleware.ValidateRequest(&controllers.RegisterRequest{}),
			authController.Register)
		public.POST("/auth/login",
			middleware.ValidateRequest(&controllers.LoginRequest{}),
			authController.Login)

		// Flight search (public)
		public.POST("/flights/search",
			middleware.ValidateRequest(&controllers.SearchFlightsRequest{}),
			flightController.SearchFlights)
	}

	// Protected routes (require authentication)
	protected := router.Group("/api")
	protected.Use(middleware.JWTMiddleware())
	{
		// User profile
		protected.GET("/profile", authController.GetProfile)
		protected.POST("/auth/refresh", middleware.RefreshToken)

		// Bookings
		protected.POST("/bookings",
			middleware.ValidateRequest(&controllers.CreateBookingRequest{}),
			bookingController.CreateBooking)
		protected.GET("/bookings", bookingController.GetUserBookings)
		protected.GET("/bookings/:id", bookingController.GetBooking)
		protected.DELETE("/bookings/:id", bookingController.CancelBooking)
	}

	// Admin routes (require admin role)
	admin := router.Group("/api/admin")
	admin.Use(middleware.JWTMiddleware(), middleware.AdminMiddleware())
	{
		// Flight management
		admin.POST("/flights",
			middleware.ValidateRequest(&controllers.CreateFlightRequest{}),
			flightController.CreateFlight)
		admin.PUT("/flights/:id",
			middleware.ValidateRequest(&controllers.CreateFlightRequest{}),
			flightController.UpdateFlight)
		admin.DELETE("/flights/:id", flightController.DeleteFlight)

		// Booking management
		admin.GET("/bookings", bookingController.GetAllBookings)
		admin.PUT("/bookings/:id/status",
			middleware.ValidateRequest(&controllers.UpdateBookingStatusRequest{}),
			bookingController.UpdateBookingStatus)
	}
}
