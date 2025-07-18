// Package models defines the Ticket data structure for flight reservations
// in the MindenAirport system.
package models

import "time"

// Ticket represents a flight reservation and booking in the airport system.
// Contains all information related to a passenger's flight booking including
// seat assignment, travel class, pricing, and booking details.
type Ticket struct {
	ID            string    `json:"id"`                      // Unique identifier for the ticket
	AirportUserID string    `json:"airportUserId"`           // ID of the passenger who booked the ticket
	Flight        string    `json:"flight"`                  // Flight ID for this ticket
	SeatNumber    string    `json:"seatNumber,omitempty"`    // Assigned seat number (e.g., "12A")
	TravelClass   string    `json:"travelClass,omitempty"`   // Travel class (Economy, Business, First)
	Price         float64   `json:"price,omitempty"`         // Ticket price in USD
	BookingDate   time.Time `json:"bookingDate,omitempty"`   // Date and time when ticket was booked
	Status        string    `json:"status,omitempty"`        // Ticket status (CONFIRMED, CANCELLED, CHECKED_IN)
	From          string    `json:"from,omitempty"`          // Origin airport code
	To            string    `json:"to,omitempty"`            // Destination airport code
	Gate          string    `json:"gate,omitempty"`          // Departure gate assignment
	BaggageClaim  string    `json:"baggageClaim,omitempty"`  // Baggage claim area for arrival
	DepartureTime string    `json:"departureTime,omitempty"` // Scheduled departure time
}
