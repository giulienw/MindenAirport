// Package models defines the Airline data structure for airline information
// management in the MindenAirport system.
package models

// Airline represents an airline company in the airport system.
// Contains essential information about airlines operating at the airport.
type Airline struct {
	ID      string `json:"id"`      // IATA airline code (e.g., "AA", "UA", "DL")
	Name    string `json:"name"`    // Full airline name (e.g., "American Airlines")
	Country string `json:"country"` // Country where airline is based
	Logo    string `json:"logo"`    // URL or path to airline logo image
	Active  bool   `json:"active"`  // Whether airline is currently active/operational
}
