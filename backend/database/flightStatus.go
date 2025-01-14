package database

import (
	"log"
	"mindenairport/models"
)

func (db Database) GetFlightStatusByID(id string) models.FlightStatus {
	var flightStatus models.FlightStatus

	err := db.QueryRow("SELECT * FROM flight_status WHERE id = :1", id).Scan(&flightStatus.ID, &flightStatus.Name)
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}

	return flightStatus
}

func (db Database) GetFlightStatuses() []models.FlightStatus {
	var flightStatuses []models.FlightStatus

	rows, err := db.Query("SELECT ID,NAME,DESCRIPTION FROM flight_status")
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var flightStatus models.FlightStatus
		err := rows.Scan(&flightStatus.ID, &flightStatus.Name, &flightStatus.Description)
		if err != nil {
			log.Fatal("Error scanning the database:", err)
		}
		flightStatuses = append(flightStatuses, flightStatus)
	}

	return flightStatuses
}
