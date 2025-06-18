package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=Chirag123 dbname=new_flights_db port=5432 sslmode=disable"
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto Migrate the models with debug logging
	log.Println("Running GORM AutoMigrate...")
	if migrateErr := db.Debug().AutoMigrate(&User{}, &Flight{}, &Booking{}, &Traveller{}); migrateErr != nil {
		log.Fatalf("Failed to auto migrate database: %v", migrateErr)
	}
	log.Println("GORM AutoMigrate completed.")
}

func main() {
	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	corsOrigin := os.Getenv("CORS_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "http://localhost:3000"
	}
	config.AllowOrigins = []string{corsOrigin}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(config))

	// Routes
	api := r.Group("/api")
	{
		// Auth routes
		api.POST("/register", registerHandler)
		api.POST("/login", loginHandler)

		// Protected routes
		protected := api.Group("/")
		protected.Use(authMiddleware())
		{
			// User routes
			protected.GET("/profile", getProfileHandler)

			// Admin routes
			admin := protected.Group("/admin")
			admin.Use(adminMiddleware())
			{
				admin.POST("/flights", createFlightHandler)
				admin.DELETE("/flights/:id", deleteFlightHandler)
				admin.PUT("/flights/:id", updateFlightHandler)
				admin.GET("/flights/:id", getFlightHandler)
			}

			// Flight routes
			protected.GET("/flights", getFlightsHandler)
			protected.POST("/bookings", createBookingHandler)
			protected.GET("/bookings", getBookingsHandler)
			protected.DELETE("/bookings/:id", deleteBookingHandler)
			protected.PUT("/bookings/:id", updateBookingHandler)

			// Location routes
			protected.GET("/locations", getUniqueLocationsHandler)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
