package controllers

import (
	"net/http"
	"time"

	"flights-project/middleware"
	"flights-project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FlightController handles flight-related requests
type FlightController struct {
	db *gorm.DB
}

// NewFlightController creates a new flight controller
func NewFlightController(db *gorm.DB) *FlightController {
	return &FlightController{db: db}
}

// CreateFlightRequest represents the request body for creating a flight
type CreateFlightRequest struct {
	FlightCode    string  `json:"flightCode" validate:"required,flightcode"`
	Origin        string  `json:"origin" validate:"required,min=3,max=3"`
	Destination   string  `json:"destination" validate:"required,min=3,max=3"`
	DepartureTime string  `json:"departureTime" validate:"required,futuredate"`
	ArrivalTime   string  `json:"arrivalTime" validate:"required,futuredate"`
	Price         float64 `json:"price" validate:"required,min=0"`
	Capacity      int     `json:"capacity" validate:"required,min=1"`
}

// CreateFlight creates a new flight
func (fc *FlightController) CreateFlight(c *gin.Context) {
	var request CreateFlightRequest
	if !middleware.GetValidatedModel(c, &request) {
		return
	}

	// Parse string dates to time.Time objects
	depTime, err := time.Parse(time.RFC3339, request.DepartureTime)
	if err != nil {
		c.Error(middleware.BadRequestError("Invalid departure time format", err))
		return
	}

	arrTime, err := time.Parse(time.RFC3339, request.ArrivalTime)
	if err != nil {
		c.Error(middleware.BadRequestError("Invalid arrival time format", err))
		return
	}

	flight := models.Flight{
		FlightCode:    request.FlightCode,
		Origin:        request.Origin,
		Destination:   request.Destination,
		DepartureTime: depTime,
		ArrivalTime:   arrTime,
		Price:         request.Price,
		Capacity:      request.Capacity,
	}

	if err := fc.db.Create(&flight).Error; err != nil {
		c.Error(middleware.InternalServerError("Failed to create flight", err))
		return
	}

	c.JSON(http.StatusCreated, flight)
}

// SearchFlightsRequest represents the request body for searching flights
type SearchFlightsRequest struct {
	Origin      string `json:"origin" validate:"required,min=3,max=3"`
	Destination string `json:"destination" validate:"required,min=3,max=3"`
	Date        string `json:"date" validate:"required,futuredate"`
}

// SearchFlights searches for flights based on criteria
func (fc *FlightController) SearchFlights(c *gin.Context) {
	var request SearchFlightsRequest
	if !middleware.GetValidatedModel(c, &request) {
		return
	}

	var flights []models.Flight
	if err := fc.db.Where("origin = ? AND destination = ? AND DATE(departure_time) = ?",
		request.Origin, request.Destination, request.Date).Find(&flights).Error; err != nil {
		c.Error(middleware.InternalServerError("Failed to search flights", err))
		return
	}

	c.JSON(http.StatusOK, flights)
}

func (fc *FlightController) GetFlights(c *gin.Context) {
	var flights []models.Flight

	// Use GORM to fetch flights with airport information
	result := fc.db.Preload("DepartureAirport").Preload("ArrivalAirport").
		Order("departure_time").
		Find(&flights)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching flights"})
		return
	}

	// Transform the data for the response
	var response []gin.H
	for _, flight := range flights {
		response = append(response, gin.H{
			"id":                flight.ID,
			"aircraft_id":       flight.AircraftID,
			"departure_airport": flight.DepartureAirport.Name,
			"arrival_airport":   flight.ArrivalAirport.Name,
			"departure_time":    flight.DepartureTime,
			"arrival_time":      flight.ArrivalTime,
			"distance":          flight.Distance,
			"status":            flight.Status,
			"price":             flight.Price,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (fc *FlightController) AddFlight(c *gin.Context) {
	var flight models.Flight
	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set timestamps
	flight.CreatedAt = time.Now()
	flight.UpdatedAt = time.Now()

	// Create flight using GORM
	result := fc.db.Create(&flight)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating flight"})
		return
	}

	c.JSON(http.StatusCreated, flight)
}

func (fc *FlightController) UpdateFlight(c *gin.Context) {
	id := c.Param("id")
	var flight models.Flight

	// First, find the flight
	result := fc.db.First(&flight, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	// Bind the new data
	// Note: When binding to an existing struct, only provided fields will be updated.
	// If a field is not provided in the JSON, its existing value in the flight struct will be retained.
	var updateRequest CreateFlightRequest // Use CreateFlightRequest to bind for consistency
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse string dates to time.Time objects for update
	depTime, err := time.Parse(time.RFC3339, updateRequest.DepartureTime)
	if err != nil {
		c.Error(middleware.BadRequestError("Invalid departure time format for update", err))
		return
	}

	arrTime, err := time.Parse(time.RFC3339, updateRequest.ArrivalTime)
	if err != nil {
		c.Error(middleware.BadRequestError("Invalid arrival time format for update", err))
		return
	}

	// Manually update fields from the request to the flight model
	flight.FlightCode = updateRequest.FlightCode
	flight.Origin = updateRequest.Origin
	flight.Destination = updateRequest.Destination
	flight.DepartureTime = depTime
	flight.ArrivalTime = arrTime
	flight.Price = updateRequest.Price
	flight.Capacity = updateRequest.Capacity

	// Update timestamp
	flight.UpdatedAt = time.Now()

	// Save the changes
	result = fc.db.Save(&flight)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating flight"})
		return
	}

	c.JSON(http.StatusOK, flight)
}

func (fc *FlightController) DeleteFlight(c *gin.Context) {
	id := c.Param("id")

	// Delete the flight using GORM
	result := fc.db.Delete(&models.Flight{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting flight"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully"})
}
