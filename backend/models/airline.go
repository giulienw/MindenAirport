package models

type Airline struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Active  bool   `json:"active"`
}
