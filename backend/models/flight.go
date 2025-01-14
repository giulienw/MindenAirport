package models

import "time"

type Flight struct {
	ID                 string    `json:"id"`
	From               string    `json:"from"`
	To                 string    `json:"to"`
	Date               time.Time `json:"date"`
	PilotID            string    `json:"pilotId"`
	PlaneID            string    `json:"planeId"`
	TerminalID         string    `json:"terminalId,omitempty"`
	StatusID           string
	Status             string     `json:"status,omitempty"`
	ScheduledDeparture time.Time  `json:"scheduledDeparture"`
	ActualDeparture    *time.Time `json:"actualDeparture,omitempty"`
	ScheduledArrival   time.Time  `json:"scheduledArrival"`
	ActualArrival      *time.Time `json:"actualArrival,omitempty"`
	Gate               string     `json:"gate,omitempty"`
	BaggageClaim       string     `json:"baggageClaim,omitempty"`
}
