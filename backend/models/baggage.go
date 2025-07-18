// Package models defines the Baggage data structure for baggage tracking
// in the MindenAirport system.
package models

// Baggage represents a piece of luggage in the airport baggage handling system.
// This model tracks baggage throughout its journey from check-in to pickup,
// including size, weight, tracking information, and current status.
type Baggage struct {
	ID              string  `json:"id"`                        // Unique identifier for the baggage item
	AirportUserID   string  `json:"airportUserId"`             // ID of the passenger who owns the baggage
	FlightID        string  `json:"flightId"`                  // ID of the flight this baggage is associated with
	Size            int     `json:"size"`                      // Size category (1=carry-on, 2=checked, 3=oversized)
	Weight          float64 `json:"weight"`                    // Weight of the baggage in pounds
	TrackingNumber  string  `json:"trackingNumber"`            // Unique tracking number for customer reference
	Status          string  `json:"status"`                    // Current status (CHECKED, IN_TRANSIT, DELIVERED, LOST)
	SpecialHandling string  `json:"specialHandling,omitempty"` // Special handling instructions (fragile, priority, etc.)
}
