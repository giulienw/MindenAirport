// Package models defines the Airport data structure for airport information
// management in the MindenAirport system.
package models

// Airport represents an airport entity in the system.
// Contains comprehensive information about airports including location,
// facilities, and operational details.
type Airport struct {
	ID               string  `json:"id"`                          // IATA airport code (e.g., "LAX", "JFK")
	Name             string  `json:"name,omitempty"`              // Full airport name
	Country          string  `json:"country"`                     // Country where airport is located
	City             string  `json:"city"`                        // City where airport is located
	Timezone         string  `json:"timezone,omitempty"`          // Airport timezone (e.g., "America/New_York")
	Elevation        float64 `json:"elevation,omitempty"`         // Airport elevation above sea level (feet)
	NumberOfTerminal int     `json:"numberOfTerminals,omitempty"` // Total number of terminals
	Latitude         float64 `json:"latitude,omitempty"`          // GPS latitude coordinate
	Longitude        float64 `json:"longitude,omitempty"`         // GPS longitude coordinate
}
