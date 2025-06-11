package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/yourusername/flights-project/models"

	"github.com/gin-gonic/gin"
)

type FlightController struct {
	DB *sql.DB
}

func NewFlightController(db *sql.DB) *FlightController {
	return &FlightController{DB: db}
}

func (fc *FlightController) GetFlights(c *gin.Context) {
	rows, err := fc.DB.Query(`
		SELECT f.id, f.aircraft_id, f.departure_airport_id, f.arrival_airport_id,
			   f.departure_time, f.arrival_time, f.distance, f.status,
			   a1.name as departure_airport, a2.name as arrival_airport
		FROM flights f
		JOIN airports a1 ON f.departure_airport_id = a1.airport_id
		JOIN airports a2 ON f.arrival_airport_id = a2.airport_id
		ORDER BY f.departure_time
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching flights"})
		return
	}
	defer rows.Close()

	var flights []gin.H
	for rows.Next() {
		var flight models.Flight
		var departureAirport, arrivalAirport string
		err := rows.Scan(
			&flight.ID, &flight.AircraftID, &flight.DepartureAirportID,
			&flight.ArrivalAirportID, &flight.DepartureTime, &flight.ArrivalTime,
			&flight.Distance, &flight.Status, &departureAirport, &arrivalAirport,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning flights"})
			return
		}

		flights = append(flights, gin.H{
			"id":                flight.ID,
			"aircraft_id":       flight.AircraftID,
			"departure_airport": departureAirport,
			"arrival_airport":   arrivalAirport,
			"departure_time":    flight.DepartureTime,
			"arrival_time":      flight.ArrivalTime,
			"distance":          flight.Distance,
			"status":            flight.Status,
		})
	}

	c.JSON(http.StatusOK, flights)
}

func (fc *FlightController) AddFlight(c *gin.Context) {
	var flight models.Flight
	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO flights (aircraft_id, departure_airport_id, arrival_airport_id,
						   departure_time, arrival_time, distance, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	err := fc.DB.QueryRow(
		query,
		flight.AircraftID,
		flight.DepartureAirportID,
		flight.ArrivalAirportID,
		flight.DepartureTime,
		flight.ArrivalTime,
		flight.Distance,
		flight.Status,
	).Scan(&flight.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating flight"})
		return
	}

	c.JSON(http.StatusCreated, flight)
}

func (fc *FlightController) UpdateFlight(c *gin.Context) {
	id := c.Param("id")
	var flight models.Flight
	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE flights
		SET aircraft_id = $1,
			departure_airport_id = $2,
			arrival_airport_id = $3,
			departure_time = $4,
			arrival_time = $5,
			distance = $6,
			status = $7
		WHERE id = $8
		RETURNING id
	`
	err := fc.DB.QueryRow(
		query,
		flight.AircraftID,
		flight.DepartureAirportID,
		flight.ArrivalAirportID,
		flight.DepartureTime,
		flight.ArrivalTime,
		flight.Distance,
		flight.Status,
		id,
	).Scan(&flight.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating flight"})
		return
	}

	c.JSON(http.StatusOK, flight)
}

func (fc *FlightController) DeleteFlight(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM flights WHERE id = $1`
	result, err := fc.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting flight"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully"})
}

func (fc *FlightController) SearchFlights(c *gin.Context) {
	// Get search parameters from query string
	departureCity := c.Query("departure_city")
	arrivalCity := c.Query("arrival_city")
	departureDate := c.Query("departure_date")
	status := c.Query("status")

	// Build the base query
	query := `
		SELECT f.id, f.aircraft_id, f.departure_airport_id, f.arrival_airport_id,
			   f.departure_time, f.arrival_time, f.distance, f.status,
			   a1.name as departure_airport, a2.name as arrival_airport,
			   a1.city as departure_city, a2.city as arrival_city
		FROM flights f
		JOIN airports a1 ON f.departure_airport_id = a1.airport_id
		JOIN airports a2 ON f.arrival_airport_id = a2.airport_id
		WHERE 1=1
	`
	args := []interface{}{}
	argCount := 1

	// Add filters based on provided parameters
	if departureCity != "" {
		query += fmt.Sprintf(" AND a1.city ILIKE $%d", argCount)
		args = append(args, "%"+departureCity+"%")
		argCount++
	}

	if arrivalCity != "" {
		query += fmt.Sprintf(" AND a2.city ILIKE $%d", argCount)
		args = append(args, "%"+arrivalCity+"%")
		argCount++
	}

	if departureDate != "" {
		query += fmt.Sprintf(" AND DATE(f.departure_time) = $%d", argCount)
		args = append(args, departureDate)
		argCount++
	}

	if status != "" {
		query += fmt.Sprintf(" AND f.status = $%d", argCount)
		args = append(args, status)
		argCount++
	}

	query += " ORDER BY f.departure_time"

	// Execute the query
	rows, err := fc.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching flights"})
		return
	}
	defer rows.Close()

	var flights []gin.H
	for rows.Next() {
		var flight models.Flight
		var departureAirport, arrivalAirport, departureCity, arrivalCity string
		err := rows.Scan(
			&flight.ID, &flight.AircraftID, &flight.DepartureAirportID,
			&flight.ArrivalAirportID, &flight.DepartureTime, &flight.ArrivalTime,
			&flight.Distance, &flight.Status, &departureAirport, &arrivalAirport,
			&departureCity, &arrivalCity,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning flights"})
			return
		}

		flights = append(flights, gin.H{
			"id":                flight.ID,
			"aircraft_id":       flight.AircraftID,
			"departure_airport": departureAirport,
			"arrival_airport":   arrivalAirport,
			"departure_city":    departureCity,
			"arrival_city":      arrivalCity,
			"departure_time":    flight.DepartureTime,
			"arrival_time":      flight.ArrivalTime,
			"distance":          flight.Distance,
			"status":            flight.Status,
		})
	}

	c.JSON(http.StatusOK, flights)
}
