package models

type Airport struct {
	ID               string  `json:"id"`
	Name             string  `json:"name,omitempty"`
	Country          string  `json:"country"`
	City             string  `json:"city"`
	Timezone         string  `json:"timezone,omitempty"`
	Elevation        float64 `json:"elevation,omitempty"`
	NumberOfTerminal int     `json:"numberOfTerminals,omitempty"`
	Latitude         float64 `json:"latitude,omitempty"`
	Longitude        float64 `json:"longitude,omitempty"`
}
