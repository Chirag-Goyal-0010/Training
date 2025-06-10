package models

type Flight struct {
	ID                 int    `json:"id"`
	AircraftID         int    `json:"aircraft_id"`
	DepartureAirportID int    `json:"departure_airport_id"`
	ArrivalAirportID   int    `json:"arrival_airport_id"`
	DepartureTime      string `json:"departure_time"`
	ArrivalTime        string `json:"arrival_time"`
	Distance           int    `json:"distance"`
	Status             string `json:"status"`
}
