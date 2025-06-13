package models

import "time"

type Flight struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	FlightCode         string    `json:"flight_code" gorm:"unique;not null"`
	Origin             string    `json:"origin" gorm:"not null"`
	Destination        string    `json:"destination" gorm:"not null"`
	AircraftID         int       `json:"aircraft_id"`
	Capacity           int       `json:"capacity" gorm:"not null"`
	DepartureAirportID int       `json:"departure_airport_id"`
	ArrivalAirportID   int       `json:"arrival_airport_id"`
	DepartureTime      time.Time `json:"departure_time"`
	ArrivalTime        time.Time `json:"arrival_time"`
	Distance           int       `json:"distance"`
	Status             string    `json:"status"`
	Price              float64   `json:"price"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`

	// Relationships
	DepartureAirport Airport `json:"departure_airport" gorm:"foreignKey:DepartureAirportID"`
	ArrivalAirport   Airport `json:"arrival_airport" gorm:"foreignKey:ArrivalAirportID"`
}

type Airport struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
