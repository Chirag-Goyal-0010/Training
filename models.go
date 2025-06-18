package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string    `gorm:"unique;not null" json:"email" binding:"required,email"`
	Password string    `json:"-" binding:"required,min=8"`
	Name     string    `json:"name" binding:"required,min=2"`
	IsAdmin  bool      `json:"is_admin"`
	Bookings []Booking `json:"bookings,omitempty"`
}

type Flight struct {
	gorm.Model
	Origin                   string    `json:"origin" binding:"required,min=3"`
	Destination              string    `json:"destination" binding:"required,min=3"`
	DepartureTime            time.Time `json:"departure_time" binding:"required"`
	ArrivalTime              time.Time `json:"arrival_time" binding:"required,gtfield=DepartureTime"`
	EconomyPrice             float64   `json:"economy_price" binding:"required,gte=0"`
	PremiumEconomyPrice      float64   `json:"premium_economy_price" binding:"required,gte=0"`
	BusinessPrice            float64   `json:"business_price" binding:"required,gte=0"`
	FirstClassPrice          float64   `json:"first_class_price" binding:"required,gte=0"`
	EconomySeats             int       `json:"economy_seats" binding:"required,gte=0"`         // Available Economy seats
	PremiumEconomySeats      int       `json:"premium_economy_seats" binding:"required,gte=0"` // Available Premium Economy seats
	BusinessSeats            int       `json:"business_seats" binding:"required,gte=0"`        // Available Business seats
	FirstClassSeats          int       `json:"first_class_seats" binding:"required,gte=0"`     // Available First Class seats
	TotalEconomySeats        int       `json:"total_economy_seats"`                            // Total Economy seats capacity
	TotalPremiumEconomySeats int       `json:"total_premium_economy_seats"`                    // Total Premium Economy seats capacity
	TotalBusinessSeats       int       `json:"total_business_seats"`                           // Total Business seats capacity
	TotalFirstClassSeats     int       `json:"total_first_class_seats"`                        // Total First Class seats capacity
	Price                    float64   `json:"price,omitempty" gorm:"-"`                       // Transient field for default price to frontend
	AvailableSeats           int       `json:"available_seats,omitempty" gorm:"-"`             // Transient field for total available seats to frontend
	IsInAir                  bool      `json:"is_in_air,omitempty" gorm:"-"`                   // Transient field to indicate if flight is in air
	IsLanded                 bool      `json:"is_landed,omitempty" gorm:"-"`                   // Transient field to indicate if flight has landed
	IsDepartingSoon          bool      `json:"is_departing_soon,omitempty" gorm:"-"`           // Transient field to indicate if flight is departing soon
	Bookings                 []Booking `json:"bookings,omitempty"`
}

type Traveller struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BookingID   uint      `json:"booking_id"`
	Title       string    `json:"title" binding:"required"`
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	DOB         time.Time `json:"dob" binding:"required"`
	Nationality string    `json:"nationality" binding:"required"`
}

type Booking struct {
	gorm.Model
	UserID      uint `json:"user_id"`
	User        User `json:"user,omitempty"`
	FlightID    uint `json:"flight_id"`
	Flight      Flight
	BookingDate time.Time   `json:"booking_date"`
	Status      string      `json:"status"`
	Seats       int         `json:"seats"`
	IsPremium   bool        `json:"is_premium"`
	TotalPrice  float64     `json:"total_price"`
	TravelClass string      `json:"travel_class" binding:"required,oneof=Economy PremiumEconomy Business FirstClass"`
	Travellers  []Traveller `json:"travellers" gorm:"constraint:OnDelete:CASCADE;"`
}
