// Package models defines the data structures used throughout the MindenAirport application.
// These models represent the core entities in the airport management system and correspond
// to database tables. All models use JSON tags for API serialization.
package models

import "time"

// TravelClass represents different classes of travel available on flights (Economy, Business, First)
type TravelClass struct {
	ID          string `json:"id"`                    // Unique identifier for the travel class
	Name        string `json:"name"`                  // Display name (e.g., "Economy", "Business")
	Description string `json:"description,omitempty"` // Optional detailed description
}

// MaintenanceLog tracks maintenance activities performed on aircraft
type MaintenanceLog struct {
	ID              string     `json:"id"`                        // Unique identifier for the maintenance record
	PlaneID         string     `json:"planeId"`                   // ID of the aircraft that was maintained
	MaintenanceDate time.Time  `json:"maintenanceDate"`           // Date when maintenance was performed
	Description     string     `json:"description"`               // Details of maintenance work performed
	Technician      string     `json:"technician"`                // Name/ID of the technician who performed work
	NextMaintenance *time.Time `json:"nextMaintenance,omitempty"` // Scheduled date for next maintenance
}

// CrewMember represents airline crew members (pilots, flight attendants, etc.)
type CrewMember struct {
	ID            string     `json:"id"`                      // Unique identifier for the crew member
	FirstName     string     `json:"firstName"`               // Crew member's first name
	LastName      string     `json:"lastName"`                // Crew member's last name
	Role          string     `json:"role"`                    // Role (e.g., "Captain", "Flight Attendant")
	LicenseNumber string     `json:"licenseNumber,omitempty"` // Professional license number
	LicenseExpiry *time.Time `json:"licenseExpiry,omitempty"` // License expiration date
}

// FlightCrew represents the assignment of crew members to specific flights
type FlightCrew struct {
	ID           string `json:"id"`           // Unique identifier for the flight crew assignment
	FlightID     string `json:"flightId"`     // ID of the flight
	CrewMemberID string `json:"crewMemberId"` // ID of the assigned crew member
	Role         string `json:"role"`         // Specific role on this flight
}

// Hangar represents aircraft storage and maintenance facilities
type Hangar struct {
	ID             string     `json:"id"`                       // Unique identifier for the hangar
	PlotID         string     `json:"plotId"`                   // ID of the plot where hangar is located
	Capacity       int        `json:"capacity,omitempty"`       // Maximum number of aircraft
	SizeSqFt       float64    `json:"sizeSqFt,omitempty"`       // Size in square feet
	Status         string     `json:"status,omitempty"`         // Current status (ACTIVE, MAINTENANCE, CLOSED)
	LastInspection *time.Time `json:"lastInspection,omitempty"` // Date of last safety inspection
	NextInspection *time.Time `json:"nextInspection,omitempty"` // Scheduled next inspection date
}

// Pilot represents licensed pilots who can operate aircraft
type Pilot struct {
	ID               string     `json:"id"`                         // Unique identifier for the pilot
	FirstName        string     `json:"firstName"`                  // Pilot's first name
	LastName         string     `json:"lastName"`                   // Pilot's last name
	LicenseType      string     `json:"licenseType,omitempty"`      // Type of pilot license (ATP, CPL, etc.)
	LicenseNumber    string     `json:"licenseNumber,omitempty"`    // License number
	LicenseExpiry    *time.Time `json:"licenseExpiry,omitempty"`    // License expiration date
	FlightHours      float64    `json:"flightHours,omitempty"`      // Total flight hours logged
	MedicalCheckDate *time.Time `json:"medicalCheckDate,omitempty"` // Last medical examination date
}

// Plane represents aircraft in the fleet
type Plane struct {
	ID                string  `json:"id"`                          // Unique identifier for the aircraft
	Name              string  `json:"name,omitempty"`              // Aircraft name/registration
	Model             string  `json:"model"`                       // Aircraft model (e.g., "Boeing 737")
	Seats             int     `json:"seats"`                       // Total passenger capacity
	AirlineID         string  `json:"airlineId,omitempty"`         // ID of the owning airline
	HangarID          string  `json:"hangarId,omitempty"`          // ID of assigned hangar
	ManufacturingYear int     `json:"manufacturingYear,omitempty"` // Year aircraft was manufactured
	MaxTakeoffWeight  float64 `json:"maxTakeoffWeight,omitempty"`  // Maximum takeoff weight in pounds
	FuelCapacity      float64 `json:"fuelCapacity,omitempty"`      // Fuel capacity in gallons
	Status            string  `json:"status,omitempty"`            // Current status (ACTIVE, MAINTENANCE, INACTIVE)
}

// Plot represents land parcels within the airport for various uses
type Plot struct {
	ID                 string     `json:"id"`       // Unique identifier for the plot
	Position           int        `json:"position"` // Position number within the airport
	TypeID             string     `json:"typeId"`   // ID referencing the plot type
	AreaSqFt           float64    `json:"areaSqFt,omitempty"`
	Status             string     `json:"status,omitempty"`
	LastMaintenance    *time.Time `json:"lastMaintenance,omitempty"`
	MaxWeightCapacity  float64    `json:"maxWeightCapacity,omitempty"`
	UtilitiesAvailable string     `json:"utilitiesAvailable,omitempty"`
}

type PlotType struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

type Shop struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	TypeID      string `json:"typeId"`
	PlotID      string `json:"plotId"`
	OpeningTime string `json:"openingTime,omitempty"`
	ClosingTime string `json:"closingTime,omitempty"`
	Description string `json:"description,omitempty"`
	IsDutyFree  bool   `json:"isDutyFree,omitempty"`
}

type ShopType struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Label         string `json:"label"`
	Category      string `json:"category,omitempty"`
	SecurityLevel string `json:"securityLevel,omitempty"`
	Description   string `json:"description,omitempty"`
	TypicalHours  string `json:"typicalHours,omitempty"`
}

type Terminal struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Capacity     int    `json:"capacity,omitempty"`
	Status       string `json:"status,omitempty"`
	FloorCount   int    `json:"floorCount,omitempty"`
	Services     string `json:"services,omitempty"`
	OpeningHours string `json:"openingHours,omitempty"`
}

type AirportUser struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Birthdate time.Time `json:"birthDate"`
	Password  string    `json:"password"`
	Active    bool      `json:"active"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Role      string    `json:"role"`
}
