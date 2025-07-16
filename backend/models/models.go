package models

import "time"

type TravelClass struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type MaintenanceLog struct {
	ID              string     `json:"id"`
	PlaneID         string     `json:"planeId"`
	MaintenanceDate time.Time  `json:"maintenanceDate"`
	Description     string     `json:"description"`
	Technician      string     `json:"technician"`
	NextMaintenance *time.Time `json:"nextMaintenance,omitempty"`
}

type CrewMember struct {
	ID            string     `json:"id"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	Role          string     `json:"role"`
	LicenseNumber string     `json:"licenseNumber,omitempty"`
	LicenseExpiry *time.Time `json:"licenseExpiry,omitempty"`
}

type FlightCrew struct {
	ID           string `json:"id"`
	FlightID     string `json:"flightId"`
	CrewMemberID string `json:"crewMemberId"`
	Role         string `json:"role"`
}

type Hangar struct {
	ID             string     `json:"id"`
	PlotID         string     `json:"plotId"`
	Capacity       int        `json:"capacity,omitempty"`
	SizeSqFt       float64    `json:"sizeSqFt,omitempty"`
	Status         string     `json:"status,omitempty"`
	LastInspection *time.Time `json:"lastInspection,omitempty"`
	NextInspection *time.Time `json:"nextInspection,omitempty"`
}

type Pilot struct {
	ID               string     `json:"id"`
	FirstName        string     `json:"firstName"`
	LastName         string     `json:"lastName"`
	LicenseType      string     `json:"licenseType,omitempty"`
	LicenseNumber    string     `json:"licenseNumber,omitempty"`
	LicenseExpiry    *time.Time `json:"licenseExpiry,omitempty"`
	FlightHours      float64    `json:"flightHours,omitempty"`
	MedicalCheckDate *time.Time `json:"medicalCheckDate,omitempty"`
}

type Plane struct {
	ID                string  `json:"id"`
	Name              string  `json:"name,omitempty"`
	Model             string  `json:"model"`
	Seats             int     `json:"seats"`
	AirlineID         string  `json:"airlineId,omitempty"`
	HangarID          string  `json:"hangarId,omitempty"`
	ManufacturingYear int     `json:"manufacturingYear,omitempty"`
	MaxTakeoffWeight  float64 `json:"maxTakeoffWeight,omitempty"`
	FuelCapacity      float64 `json:"fuelCapacity,omitempty"`
	Status            string  `json:"status,omitempty"`
}

type Plot struct {
	ID                 string     `json:"id"`
	Position           int        `json:"position"`
	TypeID             string     `json:"typeId"`
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
