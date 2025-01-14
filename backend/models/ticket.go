package models

import "time"

type Ticket struct {
	ID            string    `json:"id"`
	AirportUserID string    `json:"airportUserId"`
	Flight        string    `json:"flight"`
	SeatNumber    string    `json:"seatNumber,omitempty"`
	TravelClass   string    `json:"travelClass,omitempty"`
	Price         float64   `json:"price,omitempty"`
	BookingDate   time.Time `json:"bookingDate,omitempty"`
	Status        string    `json:"status,omitempty"`
	From          string    `json:"from,omitempty"`
	To            string    `json:"to,omitempty"`
	Gate          string    `json:"gate,omitempty"`
	BaggageClaim  string    `json:"baggageClaim,omitempty"`
	DepartureTime string    `json:"departureTime,omitempty"`
}
