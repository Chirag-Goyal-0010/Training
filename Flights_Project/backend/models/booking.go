package models

import (
	"time"
)

type Booking struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	FlightID  uint      `json:"flight_id"`
	SeatID    uint      `json:"seat_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationships
	Flight Flight `json:"flight" gorm:"foreignKey:FlightID"`
	Seat   Seat   `json:"seat" gorm:"foreignKey:SeatID"`
}

type Seat struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FlightID    uint      `json:"flight_id"`
	SeatNumber  string    `json:"seat_number"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
