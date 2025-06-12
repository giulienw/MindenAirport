package database

import (
	"fmt"
	"log"
	"mindenairport/models"
)

func (db Database) GetFlightByID(id string) models.Flight {
	var flight models.Flight

	err := db.QueryRow("SELECT * FROM flight WHERE id = :1", id).Scan(&flight.ID, &flight.From, &flight.To, &flight.Date, &flight.PilotID, &flight.PlaneID, &flight.TerminalID, &flight.StatusID, &flight.ScheduledDeparture, &flight.ActualDeparture, &flight.ScheduledArrival, &flight.ActualArrival, &flight.Gate, &flight.BaggageClaim)
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}

	return flight
}

func (db Database) GetFlights() []models.Flight {
	var flights []models.Flight

	rows, err := db.Query("SELECT * FROM flight")
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var flight models.Flight
		err := rows.Scan(&flight.ID, &flight.From, &flight.To, &flight.PilotID, &flight.PlaneID, &flight.TerminalID, &flight.StatusID, &flight.ScheduledDeparture, &flight.ActualDeparture, &flight.ScheduledArrival, &flight.ActualArrival, &flight.Gate, &flight.BaggageClaim)
		if err != nil {
			log.Fatal("Error scanning the database:", err)
		}
		flights = append(flights, flight)
	}
	fmt.Println("Retrieved flights:", len(flights))
	return flights
}

func (db Database) CreateFlight(flight models.Flight) {
	_, err := db.Exec("INSERT INTO flight (id, from, to, date, pilot_id, plane_id, terminal_id, status_id, scheduled_departure, actual_departure, scheduled_arrival, actual_arrival, gate, baggage_claim) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14)", flight.ID, flight.From, flight.To, flight.Date, flight.PilotID, flight.PlaneID, flight.TerminalID, flight.StatusID, flight.ScheduledDeparture, flight.ActualDeparture, flight.ScheduledArrival, flight.ActualArrival, flight.Gate, flight.BaggageClaim)
	if err != nil {
		log.Fatal("Error inserting into the database:", err)
	}
}

func (db Database) UpdateFlight(flight models.Flight) {
	_, err := db.Exec("UPDATE flight SET from = :1, to = :2, date = :3, pilot_id = :4, plane_id = :5, terminal_id = :6, status_id = :7, scheduled_departure = :8, actual_departure = :9, scheduled_arrival = :10, actual_arrival = :11, gate = :12, baggage_claim = :13 WHERE id = :14", flight.From, flight.To, flight.Date, flight.PilotID, flight.PlaneID, flight.TerminalID, flight.StatusID, flight.ScheduledDeparture, flight.ActualDeparture, flight.ScheduledArrival, flight.ActualArrival, flight.Gate, flight.BaggageClaim, flight.ID)
	if err != nil {
		log.Fatal("Error updating the database:", err)
	}
}

func (db Database) DeleteFlight(id string) {
	_, err := db.Exec("DELETE FROM flight WHERE id = :1", id)
	if err != nil {
		log.Fatal("Error deleting from the database:", err)
	}
}
