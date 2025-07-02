package models

type AdminDashboardStats struct {
	totalFlights    int `json:"totalFlights"`
	activeFlights   int `json:"activeFlights"`
	totalPassengers int `json:"totalPassengers"`
	totalBaggage    int `json:"totalBaggage"`
	delayedFlights  int `json:"delayedFlights"`
	lostBaggage     int `json:"lostBaggage"`
	revenue         int `json:"revenue"`
	capacity        int `json:"capacity"`
}
