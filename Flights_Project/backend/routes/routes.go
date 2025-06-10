package routes

import (
	"database/sql"

	"github.com/yourusername/flights-project/controllers"
	"github.com/yourusername/flights-project/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Initialize controllers
	authController := controllers.NewAuthController(db)
	flightController := controllers.NewFlightController(db)

	// Public routes
	router.POST("/api/auth/register", authController.Register)
	router.POST("/api/auth/login", authController.Login)

	// Protected routes
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		// Flight routes
		authorized.GET("/flights", flightController.GetFlights)

		// Admin only routes
		admin := authorized.Group("/admin")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.POST("/flights", flightController.AddFlight)
			admin.PUT("/flights/:id", flightController.UpdateFlight)
			admin.DELETE("/flights/:id", flightController.DeleteFlight)
		}
	}
}
