package database

import (
	"database/sql"
	"log"
	"mindenairport/models"
)

func (db Database) GetFlightStatuses() []models.FlightStatus {
	query := `BEGIN GetFlightStatuses(:1); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, sql.Out{Dest: &cursor})
	if err != nil {
		log.Fatal("Error calling stored procedure:", err)
	}
	defer cursor.Close()

	var flightStatuses []models.FlightStatus

	for cursor.Next() {
		var flightStatus models.FlightStatus
		err := cursor.Scan(&flightStatus.ID, &flightStatus.Name, &flightStatus.Description)
		if err != nil {
			log.Fatal("Error scanning flight status data:", err)
		}
		flightStatuses = append(flightStatuses, flightStatus)
	}

	return flightStatuses
}

func (db Database) GetFlightStatusByID(id int) models.FlightStatus {
	query := `BEGIN GetFlightStatusByID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, id, sql.Out{Dest: &cursor})
	if err != nil {
		log.Fatal("Error calling stored procedure:", err)
	}
	defer cursor.Close()

	var flightStatus models.FlightStatus

	if cursor.Next() {
		err := cursor.Scan(&flightStatus.ID, &flightStatus.Name, &flightStatus.Description)
		if err != nil {
			log.Fatal("Error scanning flight status data:", err)
		}
	}

	return flightStatus
}
