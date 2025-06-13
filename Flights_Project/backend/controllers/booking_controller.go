package controllers

import (
	"net/http"
	"time"

	"flights-project/middleware"
	"flights-project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookingController struct {
	db *gorm.DB
}

func NewBookingController(db *gorm.DB) *BookingController {
	return &BookingController{db: db}
}

type CreateBookingRequest struct {
	FlightID uint `json:"flight_id" binding:"required"`
	SeatID   uint `json:"seat_id" binding:"required"`
}

// UpdateBookingStatusRequest represents the request body for updating booking status
type UpdateBookingStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=confirmed cancelled"`
}

func (bc *BookingController) CreateBooking(c *gin.Context) {
	var req CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Check if flight exists and is available
	var flight models.Flight
	if err := bc.db.First(&flight, req.FlightID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	// Check if seat is available
	var seat models.Seat
	if err := bc.db.Where("id = ? AND flight_id = ? AND is_available = ?", req.SeatID, req.FlightID, true).First(&seat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Seat not available"})
		return
	}

	// Create booking
	booking := models.Booking{
		UserID:    userID.(uint),
		FlightID:  req.FlightID,
		SeatID:    req.SeatID,
		Status:    "confirmed",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := bc.db.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	// Update seat availability
	seat.IsAvailable = false
	if err := bc.db.Save(&seat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update seat availability"})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

func (bc *BookingController) GetUserBookings(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var bookings []models.Booking
	if err := bc.db.Where("user_id = ?", userID).Preload("Flight").Preload("Seat").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func (bc *BookingController) CancelBooking(c *gin.Context) {
	bookingID := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var booking models.Booking
	if err := bc.db.Where("id = ? AND user_id = ?", bookingID, userID).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Update booking status
	booking.Status = "cancelled"
	booking.UpdatedAt = time.Now()
	if err := bc.db.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
		return
	}

	// Update seat availability
	var seat models.Seat
	if err := bc.db.First(&seat, booking.SeatID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update seat availability"})
		return
	}

	seat.IsAvailable = true
	if err := bc.db.Save(&seat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update seat availability"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking cancelled successfully"})
}

// GetBooking retrieves a single booking by ID
func (bc *BookingController) GetBooking(c *gin.Context) {
	bookingID := c.Param("id")

	var booking models.Booking
	if err := bc.db.Preload("User").Preload("Flight").Preload("Seat").First(&booking, bookingID).Error; err != nil {
		c.Error(middleware.NotFoundError("Booking not found", err))
		return
	}

	c.JSON(http.StatusOK, booking)
}

// GetAllBookings retrieves all bookings (admin only)
func (bc *BookingController) GetAllBookings(c *gin.Context) {
	var bookings []models.Booking

	// Admin can view all bookings
	if err := bc.db.Preload("User").Preload("Flight").Preload("Seat").Find(&bookings).Error; err != nil {
		c.Error(middleware.InternalServerError("Failed to fetch all bookings", err))
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// UpdateBookingStatus updates the status of a booking (admin only)
func (bc *BookingController) UpdateBookingStatus(c *gin.Context) {
	bookingID := c.Param("id")
	var request UpdateBookingStatusRequest

	if !middleware.GetValidatedModel(c, &request) {
		return
	}

	var booking models.Booking
	if err := bc.db.First(&booking, bookingID).Error; err != nil {
		c.Error(middleware.NotFoundError("Booking not found", err))
		return
	}

	// Update booking status
	booking.Status = request.Status
	booking.UpdatedAt = time.Now()

	if err := bc.db.Save(&booking).Error; err != nil {
		c.Error(middleware.InternalServerError("Failed to update booking status", err))
		return
	}

	c.JSON(http.StatusOK, booking)
}
