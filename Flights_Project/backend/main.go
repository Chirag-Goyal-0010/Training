package main

import (
	"flights-project/config"
	"flights-project/logger"
	"flights-project/middleware"
	"flights-project/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logger.Warn("Error loading .env file", zap.Error(err))
	}
	logger.Info("Environment variables loaded")

	// Initialize logger
	logger.InitLogger()
	logger.Info("Starting application")

	// Initialize validator
	middleware.InitValidator()
	logger.Info("Validator initialized")

	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	logger.Info("Database connection established")

	// Initialize router
	router := gin.Default()

	// Setup middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorHandler())
	router.Use(logger.RequestLogger())
	router.Use(logger.ErrorLogger())

	// Setup routes
	routes.SetupRoutes(router, db)
	logger.Info("Routes configured")

	// Start server
	logger.Info("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
