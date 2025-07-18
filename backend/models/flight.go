// Package models defines the Flight data structure and related types
// for the MindenAirport flight management system.
package models

import "time"

// Flight represents a scheduled flight in the airport system.
// This is the core entity for tracking flight operations, containing
// all essential information about departure, arrival, crew, and aircraft.
type Flight struct {
	ID                 string     `json:"id"`                        // Unique flight identifier
	From               string     `json:"from"`                      // Origin airport code (IATA)
	To                 string     `json:"to"`                        // Destination airport code (IATA)
	PilotID            string     `json:"pilotId"`                   // ID of the assigned pilot
	PlaneID            string     `json:"planeId"`                   // ID of the assigned aircraft
	TerminalID         string     `json:"terminalId"`                // ID of the departure terminal
	StatusID           int        `json:"statusId"`                  // Current flight status (references FlightStatus)
	ScheduledDeparture time.Time  `json:"scheduledDeparture"`        // Planned departure time
	ActualDeparture    *time.Time `json:"actualDeparture,omitempty"` // Actual departure time (if departed)
	ScheduledArrival   time.Time  `json:"scheduledArrival"`          // Planned arrival time
	ActualArrival      *time.Time `json:"actualArrival,omitempty"`   // Actual arrival time (if arrived)
	Gate               string     `json:"gate,omitempty"`            // Assigned departure gate
	BaggageClaim       string     `json:"baggageClaim,omitempty"`    // Baggage claim area for arrival
}
