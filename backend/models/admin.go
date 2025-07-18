// Package models defines administrative data structures for the MindenAirport system.
// These models are used for administrative dashboard and reporting functions.
package models

// AdminDashboardStats represents statistical data for the administrative dashboard.
// This structure contains key metrics and KPIs for airport operations management.
// Used to provide administrators with an overview of current airport status.
type AdminDashboardStats struct {
	totalFlights    int `json:"totalFlights"`    // Total number of flights in the system
	activeFlights   int `json:"activeFlights"`   // Number of currently active/in-progress flights
	totalPassengers int `json:"totalPassengers"` // Total number of passengers processed
	totalBaggage    int `json:"totalBaggage"`    // Total number of baggage items tracked
	delayedFlights  int `json:"delayedFlights"`  // Number of delayed flights
	lostBaggage     int `json:"lostBaggage"`     // Number of lost baggage items
	revenue         int `json:"revenue"`         // Total revenue generated (in cents/smallest currency unit)
	capacity        int `json:"capacity"`        // Current airport capacity utilization
}
