package models

type Baggage struct {
	ID              string  `json:"id"`
	AirportUserID   string  `json:"airportUserId"`
	FlightID        string  `json:"flightId"`
	Size            int     `json:"size"`
	Weight          float64 `json:"weight"`
	TrackingNumber  string  `json:"trackingNumber"`
	Status          string  `json:"status"`
	SpecialHandling string  `json:"specialHandling,omitempty"`
}
