package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required,min=2"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type TravellerInput struct {
	Title       string    `json:"title" binding:"required"`
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	DOB         time.Time `json:"dob" binding:"required"`
	Nationality string    `json:"nationality" binding:"required"`
}

type CreateBookingInput struct {
	FlightID    uint             `json:"flight_id" binding:"required"`
	Seats       int              `json:"seats" binding:"required,gt=0"`
	TravelClass string           `json:"travel_class" binding:"required,oneof=Economy PremiumEconomy Business FirstClass"`
	Travellers  []TravellerInput `json:"travellers" binding:"required,dive,required"`
}

// FlightDetailsResponse is used to return detailed flight information, including booked seats per class.
type FlightDetailsResponse struct {
	Flight
	BookedEconomySeats           int `json:"booked_economy_seats"`
	BookedPremiumEconomySeats    int `json:"booked_premium_economy_seats"`
	BookedBusinessSeats          int `json:"booked_business_seats"`
	BookedFirstClassSeats        int `json:"booked_first_class_seats"`
	AvailableEconomySeats        int `json:"available_economy_seats"`
	AvailablePremiumEconomySeats int `json:"available_premium_economy_seats"`
	AvailableBusinessSeats       int `json:"available_business_seats"`
	AvailableFirstClassSeats     int `json:"available_first_class_seats"`
	TotalFlightSeats             int `json:"total_flight_seats"` // Overall total seats for the flight
}

func registerHandler(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process registration"})
		return
	}

	user := User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Name:     input.Name,
		IsAdmin:  false,
	}

	if err := db.Create(&user).Error; err != nil {
		log.Printf("Failed to create user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func loginHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	var user User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		log.Printf("Login attempt failed for email %s: %v", input.Email, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Printf("Invalid password for user %s", input.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key" // Fallback for development
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": user})
}

func getProfileHandler(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user User
	if err := db.First(&user, userID).Error; err != nil {
		log.Printf("Failed to fetch user profile for ID %v: %v", userID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func createFlightHandler(c *gin.Context) {
	var flight Flight
	if err := c.ShouldBindJSON(&flight); err != nil {
		log.Printf("Error binding JSON for flight: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Initialize Total seats with the initial available seats
	flight.TotalEconomySeats = flight.EconomySeats
	flight.TotalPremiumEconomySeats = flight.PremiumEconomySeats
	flight.TotalBusinessSeats = flight.BusinessSeats
	flight.TotalFirstClassSeats = flight.FirstClassSeats

	// Validate that at least one seat class has a positive number of seats
	if flight.EconomySeats <= 0 && flight.PremiumEconomySeats <= 0 && flight.BusinessSeats <= 0 && flight.FirstClassSeats <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one seat class must have available seats"})
		return
	}

	// Validate that prices are non-negative
	if flight.EconomyPrice < 0 || flight.PremiumEconomyPrice < 0 || flight.BusinessPrice < 0 || flight.FirstClassPrice < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Prices cannot be negative"})
		return
	}

	log.Printf("Attempting to create flight: %+v", flight) // Log the flight object received

	if err := db.Create(&flight).Error; err != nil {
		log.Printf("Error creating flight in DB: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flight"})
		return
	}

	c.JSON(http.StatusCreated, flight)
}

func deleteFlightHandler(c *gin.Context) {
	id := c.Param("id")

	if err := db.Delete(&Flight{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete flight"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully"})
}

// getFlightHandler fetches a single flight by its ID
func getFlightHandler(c *gin.Context) {
	id := c.Param("id") // Get flight ID from URL parameter

	var flight Flight
	// Preload Bookings and their Travellers
	if err := db.Preload("Bookings.Travellers").First(&flight, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
			return
		}
		log.Printf("Error fetching flight %s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flight details"})
		return
	}

	// Calculate total available seats and set economy price as default price for frontend display
	flight.Price = flight.EconomyPrice // Set EconomyPrice as the default price for display

	// Calculate booked seats per class
	var bookedSeats []struct {
		TravelClass string
		TotalSeats  int
	}
	if err := db.Model(&Booking{}).Where("flight_id = ?", id).Select("travel_class, SUM(seats) as total_seats").Group("travel_class").Scan(&bookedSeats).Error; err != nil {
		log.Printf("Error fetching booked seats for flight %s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booked seat details"})
		return
	}

	bookedEconomySeats := 0
	bookedPremiumEconomySeats := 0
	bookedBusinessSeats := 0
	bookedFirstClassSeats := 0

	for _, bs := range bookedSeats {
		switch bs.TravelClass {
		case "Economy":
			bookedEconomySeats = bs.TotalSeats
		case "PremiumEconomy":
			bookedPremiumEconomySeats = bs.TotalSeats
		case "Business":
			bookedBusinessSeats = bs.TotalSeats
		case "FirstClass":
			bookedFirstClassSeats = bs.TotalSeats
		}
	}

	totalFlightSeats := flight.TotalEconomySeats + flight.TotalPremiumEconomySeats + flight.TotalBusinessSeats + flight.TotalFirstClassSeats

	response := FlightDetailsResponse{
		Flight:                       flight,
		BookedEconomySeats:           bookedEconomySeats,
		BookedPremiumEconomySeats:    bookedPremiumEconomySeats,
		BookedBusinessSeats:          bookedBusinessSeats,
		BookedFirstClassSeats:        bookedFirstClassSeats,
		AvailableEconomySeats:        flight.EconomySeats,
		AvailablePremiumEconomySeats: flight.PremiumEconomySeats,
		AvailableBusinessSeats:       flight.BusinessSeats,
		AvailableFirstClassSeats:     flight.FirstClassSeats,
		TotalFlightSeats:             totalFlightSeats,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func getFlightsHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var flights []Flight
	var total int64

	dbQuery := db.Model(&Flight{})

	// Apply filters if provided
	origin := c.Query("origin")
	if origin != "" {
		dbQuery = dbQuery.Where("origin ILIKE ?", "%"+origin+"%")
	}

	destination := c.Query("destination")
	if destination != "" {
		dbQuery = dbQuery.Where("destination ILIKE ?", "%"+destination+"%")
	}

	status := c.Query("status")
	if status != "" {
		switch status {
		case "In Air":
			dbQuery = dbQuery.Where("departure_time < NOW() AT TIME ZONE 'UTC' AND arrival_time > NOW() AT TIME ZONE 'UTC'")
		case "Landed":
			// A flight is 'Landed' if its arrival time is in the past, AND it's not currently 'In Air' AND not 'Departing Soon'
			dbQuery = dbQuery.Where("arrival_time < NOW() AT TIME ZONE 'UTC'")
			dbQuery = dbQuery.Where("NOT (departure_time < NOW() AT TIME ZONE 'UTC' AND arrival_time > NOW() AT TIME ZONE 'UTC')")                              // Not In Air
			dbQuery = dbQuery.Where("NOT (departure_time > NOW() AT TIME ZONE 'UTC' AND (departure_time - NOW() AT TIME ZONE 'UTC') <= INTERVAL '10 minutes')") // Not Departing Soon
		case "Departing Soon":
			dbQuery = dbQuery.Where("departure_time > NOW() AT TIME ZONE 'UTC' AND (departure_time - NOW() AT TIME ZONE 'UTC') <= INTERVAL '10 minutes'")
		case "Scheduled":
			// A flight is 'Scheduled' if its departure time is in the future AND more than 10 minutes away
			dbQuery = dbQuery.Where("departure_time > NOW() AT TIME ZONE 'UTC' AND (departure_time - NOW() AT TIME ZONE 'UTC') > INTERVAL '10 minutes'")
		}
	}

	departureDateStr := c.Query("departure_date")
	if departureDateStr != "" {
		// Parse date in YYYY-MM-DD format
		departureDate, err := time.Parse("2006-01-02", departureDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid departure_date format. Expected YYYY-MM-DD."})
			return
		}
		// Filter for flights on the same day as departureDate
		dbQuery = dbQuery.Where("DATE(departure_time) = ?", departureDate.Format("2006-01-02"))
	}

	if err := dbQuery.Count(&total).Error; err != nil {
		log.Printf("Failed to count flights: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flights"})
		return
	}

	allFlights := c.Query("all_flights")

	// Only show flights with departure_time in the future for user-facing requests
	if allFlights != "true" {
		dbQuery = dbQuery.Where("departure_time > NOW() AT TIME ZONE 'UTC'")
	}

	// Order by ID in descending order to show newest flights first
	if allFlights == "true" {
		if err := dbQuery.Order("id DESC").Find(&flights).Error; err != nil {
			log.Printf("Failed to fetch all flights: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all flights"})
			return
		}
	} else {
		if err := dbQuery.Offset(offset).Limit(limit).Order("id DESC").Find(&flights).Error; err != nil {
			log.Printf("Failed to fetch flights: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flights"})
			return
		}
	}

	travelClass := c.Query("travel_class")

	// Calculate total available seats and set economy price as default price for frontend display
	for i := range flights {
		switch travelClass {
		case "Economy":
			flights[i].Price = flights[i].EconomyPrice
		case "PremiumEconomy":
			flights[i].Price = flights[i].PremiumEconomyPrice
		case "Business":
			flights[i].Price = flights[i].BusinessPrice
		case "FirstClass":
			flights[i].Price = flights[i].FirstClassPrice
		default:
			flights[i].Price = flights[i].EconomyPrice // Default to Economy Price
		}
		flights[i].AvailableSeats = flights[i].EconomySeats + flights[i].PremiumEconomySeats + flights[i].BusinessSeats + flights[i].FirstClassSeats

		// Calculate real-time flight status (these are for display only, filtering is done in SQL)
		now := time.Now()
		flights[i].IsInAir = flights[i].DepartureTime.Before(now) && flights[i].ArrivalTime.After(now)
		flights[i].IsLanded = flights[i].ArrivalTime.Before(now)
		flights[i].IsDepartingSoon = flights[i].DepartureTime.After(now) && flights[i].DepartureTime.Sub(now) <= 10*time.Minute
	}

	c.JSON(http.StatusOK, gin.H{
		"data": flights,
		"meta": gin.H{
			"total":  total,
			"page":   page,
			"limit":  limit,
			"offset": offset,
		},
	})
}

func createBookingHandler(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input CreateBookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if len(input.Travellers) != input.Seats {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Number of travellers must match number of seats."})
		return
	}

	for i, t := range input.Travellers {
		if t.Title == "" || t.FirstName == "" || t.LastName == "" || t.Nationality == "" || t.DOB.IsZero() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "All traveller details are required for traveller #" + strconv.Itoa(i+1)})
			return
		}
	}

	var bookingResponse struct {
		IsPremium  bool    `json:"is_premium"`
		TotalPrice float64 `json:"total_price"`
	}

	// Use transaction for booking creation
	err := db.Transaction(func(tx *gorm.DB) error {
		// Check if flight exists
		var flight Flight
		if err := tx.First(&flight, input.FlightID).Error; err != nil {
			return err
		}

		// Determine base price based on travel class
		var basePricePerSeat float64
		switch input.TravelClass {
		case "Economy":
			basePricePerSeat = flight.EconomyPrice
		case "PremiumEconomy":
			basePricePerSeat = flight.PremiumEconomyPrice
		case "Business":
			basePricePerSeat = flight.BusinessPrice
		case "FirstClass":
			basePricePerSeat = flight.FirstClassPrice
		default:
			return fmt.Errorf("invalid travel class")
		}

		// Check flight departure time
		timeUntilDeparture := time.Until(flight.DepartureTime)

		// If less than 15 minutes until departure, booking is not allowed
		if timeUntilDeparture < 15*time.Minute {
			return fmt.Errorf("cannot book flight: departure is less than 15 minutes away")
		}

		// Determine if this is a premium booking (between 60 and 15 minutes before departure)
		isPremium := timeUntilDeparture < 60*time.Minute && timeUntilDeparture >= 15*time.Minute

		// Calculate total price
		calculatedPrice := basePricePerSeat * float64(input.Seats)
		totalPrice := calculatedPrice
		if isPremium {
			totalPrice = calculatedPrice * 1.30 // 30% extra charge for premium booking
		}

		// Store response values
		bookingResponse.IsPremium = isPremium
		bookingResponse.TotalPrice = totalPrice

		// Check if enough seats are available for the selected travel class
		switch input.TravelClass {
		case "Economy":
			if input.Seats > flight.EconomySeats {
				return fmt.Errorf("not enough economy seats available")
			}
			flight.EconomySeats -= input.Seats
		case "PremiumEconomy":
			if input.Seats > flight.PremiumEconomySeats {
				return fmt.Errorf("not enough premium economy seats available")
			}
			flight.PremiumEconomySeats -= input.Seats
		case "Business":
			if input.Seats > flight.BusinessSeats {
				return fmt.Errorf("not enough business seats available")
			}
			flight.BusinessSeats -= input.Seats
		case "FirstClass":
			if input.Seats > flight.FirstClassSeats {
				return fmt.Errorf("not enough first class seats available")
			}
			flight.FirstClassSeats -= input.Seats
		default:
			return fmt.Errorf("invalid travel class")
		}

		booking := Booking{
			UserID:      userID.(uint),
			FlightID:    input.FlightID,
			BookingDate: time.Now(),
			Status:      "Confirmed",
			Seats:       input.Seats,
			IsPremium:   isPremium,
			TotalPrice:  totalPrice,
			TravelClass: input.TravelClass,
		}

		if err := tx.Create(&booking).Error; err != nil {
			return err
		}

		// Save travellers
		for _, t := range input.Travellers {
			trav := Traveller{
				BookingID:   booking.ID,
				Title:       t.Title,
				FirstName:   t.FirstName,
				LastName:    t.LastName,
				DOB:         t.DOB,
				Nationality: t.Nationality,
			}
			if err := tx.Create(&trav).Error; err != nil {
				return err
			}
		}

		// Update flight seats
		if err := tx.Save(&flight).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("Failed to create booking: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Booking created successfully",
		"is_premium":  bookingResponse.IsPremium,
		"total_price": bookingResponse.TotalPrice,
	})
}

func getBookingsHandler(c *gin.Context) {
	userID, _ := c.Get("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var bookings []Booking
	var total int64

	if err := db.Model(&Booking{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		log.Printf("Failed to count bookings: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	if err := db.Preload("Flight").Offset(offset).Limit(limit).Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		log.Printf("Failed to fetch bookings: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookings,
		"meta": gin.H{
			"total":  total,
			"page":   page,
			"limit":  limit,
			"offset": offset,
		},
	})
}

func deleteBookingHandler(c *gin.Context) {
	id := c.Param("id")

	err := db.Transaction(func(tx *gorm.DB) error {
		var booking Booking
		if err := tx.First(&booking, id).Error; err != nil {
			return fmt.Errorf("booking not found")
		}

		// Restore seats to the flight
		var flight Flight
		if err := tx.First(&flight, booking.FlightID).Error; err != nil {
			return fmt.Errorf("flight not found")
		}

		switch booking.TravelClass {
		case "Economy":
			flight.EconomySeats += booking.Seats
		case "PremiumEconomy":
			flight.PremiumEconomySeats += booking.Seats
		case "Business":
			flight.BusinessSeats += booking.Seats
		case "FirstClass":
			flight.FirstClassSeats += booking.Seats
		}

		if err := tx.Save(&flight).Error; err != nil {
			return fmt.Errorf("failed to restore seats to flight")
		}

		if err := tx.Delete(&booking).Error; err != nil {
			return fmt.Errorf("failed to delete booking")
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully and seats restored"})
}

func updateBookingHandler(c *gin.Context) {
	id := c.Param("id")

	var input CreateBookingInput // Using CreateBookingInput for simplicity, adjust as needed
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		var booking Booking
		if err := tx.First(&booking, id).Error; err != nil {
			return fmt.Errorf("booking not found")
		}

		// Restore old seats to the flight first
		var flight Flight
		if err := tx.First(&flight, booking.FlightID).Error; err != nil {
			return fmt.Errorf("flight not found")
		}

		switch booking.TravelClass {
		case "Economy":
			flight.EconomySeats += booking.Seats
		case "PremiumEconomy":
			flight.PremiumEconomySeats += booking.Seats
		case "Business":
			flight.BusinessSeats += booking.Seats
		case "FirstClass":
			flight.FirstClassSeats += booking.Seats
		}

		// Check if enough seats are available for the new travel class and quantity
		switch input.TravelClass {
		case "Economy":
			if input.Seats > flight.EconomySeats {
				return fmt.Errorf("not enough economy seats available for update")
			}
			flight.EconomySeats -= input.Seats
		case "PremiumEconomy":
			if input.Seats > flight.PremiumEconomySeats {
				return fmt.Errorf("not enough premium economy seats available for update")
			}
			flight.PremiumEconomySeats -= input.Seats
		case "Business":
			if input.Seats > flight.BusinessSeats {
				return fmt.Errorf("not enough business seats available for update")
			}
			flight.BusinessSeats -= input.Seats
		case "FirstClass":
			if input.Seats > flight.FirstClassSeats {
				return fmt.Errorf("not enough first class seats available for update")
			}
			flight.FirstClassSeats -= input.Seats
		default:
			return fmt.Errorf("invalid travel class for update")
		}

		if err := tx.Save(&flight).Error; err != nil {
			return fmt.Errorf("failed to update flight seats during booking update")
		}

		// Update booking details
		booking.Seats = input.Seats
		booking.TravelClass = input.TravelClass

		// Recalculate price if needed (based on premium logic)
		// Note: This assumes the price is per seat and the premium logic is applied per booking
		timeUntilDeparture := time.Until(flight.DepartureTime)
		isPremium := timeUntilDeparture < 60*time.Minute && timeUntilDeparture >= 15*time.Minute
		basePrice := 0.0 // Initialize basePrice for recalculation
		switch input.TravelClass {
		case "Economy":
			basePrice = flight.EconomyPrice
		case "PremiumEconomy":
			basePrice = flight.PremiumEconomyPrice
		case "Business":
			basePrice = flight.BusinessPrice
		case "FirstClass":
			basePrice = flight.FirstClassPrice
		}
		calculatedPrice := basePrice * float64(input.Seats)
		totalPrice := calculatedPrice
		if isPremium {
			totalPrice = calculatedPrice * 1.30
		}
		booking.IsPremium = isPremium
		booking.TotalPrice = totalPrice

		if err := tx.Save(&booking).Error; err != nil {
			return fmt.Errorf("failed to update booking details")
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking updated successfully"})
}

func updateFlightHandler(c *gin.Context) {
	id := c.Param("id")
	var flight Flight

	if err := db.First(&flight, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	var updatedFlight Flight
	if err := c.ShouldBindJSON(&updatedFlight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Validate that at least one seat class has a positive number of seats
	if updatedFlight.EconomySeats <= 0 && updatedFlight.PremiumEconomySeats <= 0 && updatedFlight.BusinessSeats <= 0 && updatedFlight.FirstClassSeats <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one seat class must have available seats"})
		return
	}

	// Validate that prices are non-negative
	if updatedFlight.EconomyPrice < 0 || updatedFlight.PremiumEconomyPrice < 0 || updatedFlight.BusinessPrice < 0 || updatedFlight.FirstClassPrice < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Prices cannot be negative"})
		return
	}

	// Update only the fields that are allowed to be updated
	flight.Origin = updatedFlight.Origin
	flight.Destination = updatedFlight.Destination
	flight.DepartureTime = updatedFlight.DepartureTime
	flight.ArrivalTime = updatedFlight.ArrivalTime
	flight.EconomyPrice = updatedFlight.EconomyPrice
	flight.PremiumEconomyPrice = updatedFlight.PremiumEconomyPrice
	flight.BusinessPrice = updatedFlight.BusinessPrice
	flight.FirstClassPrice = updatedFlight.FirstClassPrice
	flight.EconomySeats = updatedFlight.EconomySeats
	flight.PremiumEconomySeats = updatedFlight.PremiumEconomySeats
	flight.BusinessSeats = updatedFlight.BusinessSeats
	flight.FirstClassSeats = updatedFlight.FirstClassSeats

	if err := db.Save(&flight).Error; err != nil {
		log.Printf("Error updating flight in DB: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update flight"})
		return
	}

	c.JSON(http.StatusOK, flight)
}

func getUniqueLocationsHandler(c *gin.Context) {
	var origins []string
	// Fetch unique origin cities by joining with Airports table
	if err := db.Model(&Flight{}).Joins("left join airports on flights.departure_airport_id = airports.airport_id").Where("airports.city IS NOT NULL").Distinct("airports.city").Pluck("airports.city", &origins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch origins"})
		return
	}

	var destinations []string
	// Fetch unique destination cities by joining with Airports table
	if err := db.Model(&Flight{}).Joins("left join airports on flights.arrival_airport_id = airports.airport_id").Where("airports.city IS NOT NULL").Distinct("airports.city").Pluck("airports.city", &destinations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch destinations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"origins":      origins,
		"destinations": destinations,
	})
}
